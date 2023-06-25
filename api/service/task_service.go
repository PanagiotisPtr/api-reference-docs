package service

import "github.com/panagiotisptr/api-reference-docs/api/model"

type ErrorResponse struct {
	Error string `json:"error"`
}

type TaskService struct {
	Tasks []model.Task
}

type CreateTaskRequest struct {
	Task model.Task `json:"task"`
}

type CreateTaskResponse struct {
	Task model.Task `json:"task"`
}

type UpdateTaskRequest struct {
	UpdatedTask model.Task `json:"updatedTask"`
}

type UpdateTaskResponse struct {
	UpdatedTask model.Task `json:"updatedTask"`
}

type DeleteTaskRequest struct{}

type DeleteTaskResponse struct {
	DeletedTask model.Task `json:"deletedTask"`
}

type GetTaskRequest struct{}

type GetTaskResponse struct {
	Task model.Task `json:"task"`
}

type ListTasksRequest struct{}

type ListTasksResponse struct {
	Tasks []model.Task `json:"tasks"`
}

// CreateTask creates a new task and adds it to the task list.
func (ts *TaskService) CreateTask(req CreateTaskRequest) CreateTaskResponse {
	task := req.Task
	ts.Tasks = append(ts.Tasks, task)
	return CreateTaskResponse{Task: task}
}

// UpdateTask updates the details of an existing task.
func (ts *TaskService) UpdateTask(id int64, req UpdateTaskRequest) UpdateTaskResponse {
	updatedTask := req.UpdatedTask

	for i, task := range ts.Tasks {
		if task.ID == id {
			ts.Tasks[i] = updatedTask
			return UpdateTaskResponse{UpdatedTask: updatedTask}
		}
	}

	return UpdateTaskResponse{}
}

// DeleteTask removes a task from the task list.
func (ts *TaskService) DeleteTask(id int64, req DeleteTaskRequest) DeleteTaskResponse {
	for i, task := range ts.Tasks {
		if task.ID == id {
			ts.Tasks = append(ts.Tasks[:i], ts.Tasks[i+1:]...)
			return DeleteTaskResponse{DeletedTask: task}
		}
	}

	return DeleteTaskResponse{}
}

// GetTask retrieves a task with the given ID from the task list.
func (ts *TaskService) GetTask(id int64, req GetTaskRequest) GetTaskResponse {
	for _, task := range ts.Tasks {
		if task.ID == id {
			return GetTaskResponse{Task: task}
		}
	}

	return GetTaskResponse{}
}

// ListTasks returns all the tasks in the task list.
func (ts *TaskService) ListTasks(req ListTasksRequest) ListTasksResponse {
	return ListTasksResponse{Tasks: ts.Tasks}
}
