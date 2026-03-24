package model

import "time"

type Product struct {
	ID         int        `json:"id"`
	Name       string     `json:"name" binding:"required"`
	Price      float32    `json:"price" binding:"required"`
	CreatedAt  time.Time  `json:"created_at"`
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}
