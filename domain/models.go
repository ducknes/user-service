package domain

type User struct {
	Id                string `json:"id,omitempty"`
	Username          string `bson:"username"`
	Password          string `bson:"password"`
	Surname           string `json:"surname"`
	Name              string `json:"name"`
	Lastname          string `json:"lastname"`
	RegisteredObjects int    `json:"registeredObjects"`
	Role              Role   `json:"role"`
}
