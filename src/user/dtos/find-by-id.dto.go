package dtos

import "time"

type FindByIDDto struct {
	ID        string    `json:"Id"`
	Username  string    `json:"Username"`
	Email     string    `json:"Email"`
	Role      string    `json:"Role"`
	IsActive  bool      `json:"IsActive"`
	CreatedAt time.Time `json:"CreatedAt"`
}
