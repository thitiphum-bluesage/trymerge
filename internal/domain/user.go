package domain

// User represents the user model for our application
type User struct {
    ID       uint   `json:"id"`
    Email    string `json:"email"`
    Password string `json:"-"`
}
