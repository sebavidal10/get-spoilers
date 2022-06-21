package router

import (
	"github.com/gorilla/mux"
	"github.com/sebavidal10/get-spoilers/middleware"
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()

		router.HandleFunc("/api/spoilers", middleware.GetAllSpoiler).Methods("GET", "OPTIONS")
    // router.HandleFunc("/api/user/{id}", middleware.GetUser).Methods("GET", "OPTIONS")
    // router.HandleFunc("/api/user", middleware.GetAllUser).Methods("GET", "OPTIONS")
    // router.HandleFunc("/api/newuser", middleware.CreateUser).Methods("POST", "OPTIONS")
    // router.HandleFunc("/api/user/{id}", middleware.UpdateUser).Methods("PUT", "OPTIONS")
    // router.HandleFunc("/api/deleteuser/{id}", middleware.DeleteUser).Methods("DELETE", "OPTIONS")

    return router
}
