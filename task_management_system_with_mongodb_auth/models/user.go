package models

type User struct {
	ID       uint   `json:"id" bson:"id,omitempty"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"-" bson:"password"`
	Role    string `json:"role" bson:"role"`
  }