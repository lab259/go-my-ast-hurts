package models

type User struct {
	ID      int64  `json:"id"`
    Name    string `json:"name,lastName"`
	Address string `json:"address" bson:""`
}