package database

import "time"

type User struct {
	Id                string `bson:"_id,omitempty"`
	Username          string `bson:"username"`
	Password          string `bson:"password"`
	Surname           string `bson:"surname"`
	Name              string `bson:"name"`
	Lastname          string `bson:"lastname"`
	RegisteredObjects int    `bson:"registeredObjects"`
	Role              uint   `bson:"role"`
}

type ApproveMessage struct {
	UserId    string `json:"user_id"`
	ProductId string `json:"product_id"`
}

type ApprovedItem struct {
	ProductId   string    `json:"product_id"`
	ApproveTime time.Time `json:"approve_time"`
}
