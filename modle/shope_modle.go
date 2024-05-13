package modle

import "time"

type Shopkeeper struct {
	ShopeNo   int       `json:"shope_no"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	ShopeType string    `json:"shope_type"`
	CreatedAt time.Time `json:"created_at"`
}
