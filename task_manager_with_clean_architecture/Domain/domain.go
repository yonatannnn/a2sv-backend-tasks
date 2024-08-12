package domain

type Task struct {
	ID          int    `json: "id" bson: "_id,omitempty"`
	Title       string `json: "title" bson: "title"`
	Description string `json: "description" bson: "description"`
	Completed   bool   `json: "completed" bson: "completed"`
}

type TaskRepository interface {
	CreateTask(task Task) error
	GetTaskByID(id int) (Task, error)
	UpdateTask(task Task) error
	DeleteTask(id int) error
	GetAllTasks() ([]*Task, error)
}

type TaskUsecase interface {
	CreateTask(task Task) error
	GetTaskByID(id int) (Task, error)
	UpdateTask(task Task) error
	DeleteTask(id int) error
	GetAllTasks() ([]*Task, error)
}

type User struct {
	ID       uint   `json:"id" bson:"_id,omitempty"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     string `json:"role" bson:"role"`
}

type UserRepository interface {
	Register(user User) (User, error)
	Login(username, password string) (User, error)
	PromoteUser(userID int) error
	GetUserByID(id int) (User, error)
}
