package models

type Spoiler struct {
	ID      int64		`json:"id"`
	Content	string 	`json:"content"`
	Movie 	string 	`json:"movie"`
}
