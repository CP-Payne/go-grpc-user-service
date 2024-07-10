package model

// UserData defines a users data
type UserData struct {
	ID        int     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	City      string  `json:"city"`
	Phone     string  `json:"phone"`
	Height    float32 `json:"height"`
	Married   bool    `json:"married"`
}
