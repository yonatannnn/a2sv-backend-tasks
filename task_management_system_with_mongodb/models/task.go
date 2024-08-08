package models
type Task struct {
	ID          int `bson:"_id,omitempty"`
	Title       string             `json: "title" bson: "title"`
	Description string             `json: "description" bson: "description"`
	Completed   bool               `json: "completed" bson: "completed"`
}
