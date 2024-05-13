package modle

import "time"

type User struct {
	UserId    int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Gender    string    `json:"gender"`
	Mobile    string    `json:"mobile"`
	Adult     bool      `json:"adult"`
	CreatedAt time.Time `json:"created_at"`
}
