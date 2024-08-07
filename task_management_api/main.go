package main

import (
    "api/task_manager/router"
	"api/task_manager/data"
)
func main() {
    data.InitData()
    r := router.SetupRouter()
	r.Run("localhost:3000")
}