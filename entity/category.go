package entity

import "time"

type Categories []Category

type Category struct {
	ID        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Status    bool      `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CategoryRequest struct {
	Name   string `json:"name"`
	Status bool   `json:"status,omitempty"`
}
