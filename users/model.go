package users

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	// TODO password protection implementation
}
