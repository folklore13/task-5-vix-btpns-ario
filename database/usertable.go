package database

type UserTable struct {
	UserID string `json:"user_id"`
	UserName string `json:"user_name"`
	Email string `json:"user_email"`
	Password string `json:"user_password"`
}