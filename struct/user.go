package struct

import "time"

type User struct {
	ID          int       `json:"id"`
	NIM         int       `json:"nim"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Role        string    `json:"role"`
	LoggedIn    bool      `json:"logged_in"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}