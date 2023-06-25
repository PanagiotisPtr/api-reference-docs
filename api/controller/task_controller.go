package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/panagiotisptr/api-reference-docs/api/service"
)

type TaskController struct {
	TaskService *service.TaskService
}

// POST /tasks/create task
// Creates a new task using the provided task details.
//
// Request:
// curl -X POST -d <<< EOF
// {
//     "task": {
//         "title": "Task Title",
//         "description": "Task Description",
//         "done": false
//     }
// }
// EOF
//
// Response (200):
// {
//     "task": {
//         "id": 1,
//         "title": "Task Title",
//         "description": "Task Description",
//         "done": false
//     }
// }
//
// Response (400):
// {
//     "error": "Invalid request body"
// }
//
// Request body: {empty}
// Response 200 (application/json): service.CreateTaskResponse
// Response 400 (application/json): service.ErrorResponse
func (tc *TaskController) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	var req service.CreateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp := tc.TaskService.CreateTask(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// PUT /tasks/{id} task
// Updates the details of the task with the specified ID.
//
// Request:
// curl -X PUT -d <<< EOF
// {
//     "updatedTask": {
//         "title": "Updated Title",
//         "description": "Updated Description",
//         "done": true
//     }
// }
// EOF
//
// Response (200):
// {
//     "updatedTask": {
//         "id": 1,
//         "title": "Updated Title",
//         "description": "Updated Description",
//         "done": true
//     }
// }
//
// Response (400):
// {
//     "error": "Invalid request body"
// }
//
// Response (404):
// {
//     "error": "Task not found"
// }
//
// Request body: {empty}
// Response 200 (application/json): service.UpdateTaskResponse
// Response 404 (application/json): service.ErrorResponse
// Response 400 (application/json): service.ErrorResponse
func (tc *TaskController) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	taskID, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var req service.UpdateTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp := tc.TaskService.UpdateTask(taskID, req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// DELETE /tasks/{id} task
// Deletes the task with the specified ID.
//
// Request:
// curl -X DELETE http://localhost:8080/tasks/1
//
// Response 200 (application/json):
// {
//     "deletedTask": {
//         "id": 1,
//         "title": "Task Title",
//         "description": "Task Description",
//         "done": false
//     }
// }
//
// Response (404):
// {
//     "error": "Task not found"
// }
//
// Request body: {empty}
// Response 200 (application/json): service.DeleteTaskResponse
// Response 404 (application/json): service.ErrorResponse
func (tc *TaskController) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	taskID, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	resp := tc.TaskService.DeleteTask(taskID, service.DeleteTaskRequest{})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// GET /tasks/{id} task
// Retrieves the task with the specified ID.
//
// Request:
// curl -X GET http://localhost:8080/tasks/1
//
// Response (200):
// {
//     "task": {
//         "id": 1,
//         "title": "Task Title",
//         "description": "Task Description",
//         "done": false
//     }
// }
//
// Response (404):
// {
//     "error": "Task not found"
// }
//
// Request body: {empty}
// Response 200 (application/json): service.GetTaskResponse
// Response 400 (application/json): service.ErrorResponse
func (tc *TaskController) GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	taskID, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	resp := tc.TaskService.GetTask(taskID, service.GetTaskRequest{})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// GET /tasks/list task
// Retrieves a list of all tasks.
//
// Request:
// curl -X GET http://localhost:8080/tasks
//
// Response (200):
// {
//     "tasks": [
//         {
//             "id": 1,
//             "title": "Task 1",
//             "description": "Description 1",
//             "done": false
//         },
//         {
//             "id": 2,
//             "title": "Task 2",
//             "description": "Description 2",
//             "done": true
//         }
//     ]
// }
//
// Response (400):
// {
//     "error": "Invalid request"
// }
//
// Request body: {empty}
// Response 200 (application/json): service.ListTasksResponse
// Response 400 (application/json): service.ErrorResponse
func (tc *TaskController) ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	resp := tc.TaskService.ListTasks(service.ListTasksRequest{})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
