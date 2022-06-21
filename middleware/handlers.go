package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sebavidal10/get-spoilers/models"

	// "strconv"
	// "github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// response format
type response struct {
    ID      int64  `json:"id,omitempty"`
    Message string `json:"message,omitempty"`
}

// create connection with postgres db
func createConnection() *sql.DB {
    // load .env file
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Open the connection
    db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

    if err != nil {
        panic(err)
    }

    // check the connection
    err = db.Ping()

    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected!")
    // return the connection
    return db
}

// GetAllSpoiler will return all the spoilers
func GetAllSpoiler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    // get all the spoilers in the db
    spoilers, err := getAllSpoilers()

    if err != nil {
        log.Fatalf("Unable to get all spoiler. %v", err)
    }

    // send all the spoilers as response
    json.NewEncoder(w).Encode(spoilers)
}

func getAllSpoilers() ([]models.Spoiler, error) {
    // create the postgres db connection
    db := createConnection()

    // close the db connection
    defer db.Close()

    var spoilers []models.Spoiler

    // create the select sql query
    sqlStatement := `SELECT * FROM spoilers`

    // execute the sql statement
    rows, err := db.Query(sqlStatement)

    if err != nil {
        log.Fatalf("Unable to execute the query. %v", err)
    }

    // close the statement
    defer rows.Close()

    // iterate over the rows
    for rows.Next() {
        var spoiler models.Spoiler

        // unmarshal the row object to spoiler
        err = rows.Scan(&spoiler.ID, &spoiler.Content, &spoiler.Movie)

        if err != nil {
            log.Fatalf("Unable to scan the row. %v", err)
        }

        // append the spoiler in the spoilers slice
        spoilers = append(spoilers, spoiler)

    }

    // return empty spoiler on error
    return spoilers, err
}
