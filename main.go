package main

import (
	"fmt"
	"net/http"

	"github.com/aaron-seo/godoit/tasks"
	"github.com/aaron-seo/godoit/users"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()

	// Root route
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root"))
	})

	// tasks routes (in a subrouter that we throw away)
	router.Route("/tasks", func(r chi.Router) {
		r.Get("/", tasks.ListTasks)   // GET /tasks
		r.Post("/", tasks.CreateTask) // POST /tasks

		r.Route("/{taskID}", func(r chi.Router) {
			r.Use(tasks.TaskCtx)            // Loads a Task object into a request payload
			r.Get("/", tasks.GetTask)       // GET /tasks/123
			r.Put("/", tasks.UpdateTask)    // PUT /tasks/123
			r.Delete("/", tasks.DeleteTask) // DELETE /tasks/123
		})
	})

	// users routes
	router.Route("/users", func(r chi.Router) {
		r.Get("/", users.ListUsers)   // GET /users
		r.Post("/", users.CreateUser) // POST /users

		r.Route("/{userID}", func(r chi.Router) {
			r.Use(users.UserCtx)            // Loads a User object into a request payload
			r.Get("/", users.GetUser)       // GET /users/123
			r.Put("/", users.UpdateUser)    // PUT /users/123
			r.Delete("/", users.DeleteUser) // DELETE /users/123
		})
	})

	fmt.Println("GoDoIt is online.")
	http.ListenAndServe(":8080", router)
}
