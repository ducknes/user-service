package database

type User struct {
	Id                string `bson:"_id,omitempty"`
	Surname           string `bson:"surname"`
	Name              string `bson:"name"`
	Lastname          string `bson:"lastname"`
	RegisteredObjects int    `bson:"registeredObjects"`
	Role              uint   `bson:"role"`
}
