package users

import "net/http"

// GET /users
func ListUsers(w http.ResponseWriter, r *http.Request) {

}

// POST /users
func CreateUser(w http.ResponseWriter, r *http.Request) {

}

// GET /users/123
func GetUser(w http.ResponseWriter, r *http.Request) {

}

// PUT /users/123
func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

// DELETE /users/123
func DeleteUser(w http.ResponseWriter, r *http.Request) {

}

// Middleware to load a User object into request payload
func UserCtx(next http.Handler) http.Handler {
	return nil
}
