package models

type User struct {
	ID       uint   `json:"id" bson:"id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}
