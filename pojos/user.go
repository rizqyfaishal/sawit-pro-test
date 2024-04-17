package pojos

import "time"

type User struct {
	Id                int64     `json:"id"`
	PhoneNumber       string    `json:"phone_number"`
	FullName          string    `json:"full_name"`
	LoginSuccessCount int64     `json:"login_success_count"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
