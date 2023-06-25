package main

import (
	"log"
	"net/http"

	"github.com/panagiotisptr/api-reference-docs/api/controller"
	"github.com/panagiotisptr/api-reference-docs/api/service"
)

func main() {
	taskService := &service.TaskService{}
	taskController := &controller.TaskController{
		TaskService: taskService,
	}

	http.HandleFunc("/tasks/create", taskController.CreateTaskHandler)
	http.HandleFunc("/tasks/list", taskController.ListTasksHandler)
	http.HandleFunc("/tasks/update", taskController.UpdateTaskHandler)
	http.HandleFunc("/tasks/delete", taskController.DeleteTaskHandler)
	http.HandleFunc("/tasks/get", taskController.GetTaskHandler)

	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
