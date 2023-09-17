package entity

import "time"

type Books []Book

type Book struct {
	ID          string    `json:"id" db:"id"`
	Tittle      string    `json:"tittle" db:"tittle"`
	Description string    `json:"description" db:"description"`
	Price       int       `json:"price" db:"price"`
	Image       string    `json:"image" db:"image"`
	Categories  string    `json:"categories,omitempty" db:"categories"`
	Keywords    string    `json:"keywords,omitempty" db:"keywords"`
	Stock       int       `json:"stock,omitempty" db:"stock"`
	Publisher   string    `json:"publisher,omitempty" db:"publisher"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type BookRequest struct {
	Tittle      string `json:"tittle"`
	Description string `json:"description,omitempty"`
	Price       int    `json:"price,omitempty"`
	Image       string `json:"image,omitempty"`
	Categories  string `json:"categories,omitempty"`
	Keywords    string `json:"keywords,omitempty"`
	Stock       int    `json:"stock,omitempty"`
	Publisher   string `json:"publisher,omitempty"`
}
