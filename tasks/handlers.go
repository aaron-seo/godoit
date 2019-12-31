package tasks

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

// GET /tasks
func ListTasks(w http.ResponseWriter, r *http.Request) {

}

// POST /tasks
func CreateTask(w http.ResponseWriter, r *http.Request) {

}

// GET /tasks/123
func GetTask(w http.ResponseWriter, r *http.Request) {

}

// PUT /tasks/123
func UpdateTask(w http.ResponseWriter, r *http.Request) {

}

// DELETE /tasks/123
func DeleteTask(w http.ResponseWriter, r *http.Request) {

}

// Middleware to load a Task object into request payload
func TaskCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//var task Task

		if taskID := chi.URLParam(r, "taskID"); taskID != "" {
			fmt.Println(taskID)
		}
	})
}
