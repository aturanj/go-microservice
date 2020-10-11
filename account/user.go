package account

import "context"

// User struct with json keys
type User struct {
	ID       string `json:"id,omitemty"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Repository interface
type Repository interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, ID string) (string, error)
}
