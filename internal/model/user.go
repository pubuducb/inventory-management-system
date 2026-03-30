package model

import "time"

type User struct {
	ID         int        `json:"id"`
	Name       string     `json:"name" binding:"required,min=3,max=100"`
	Email      string     `json:"email" binding:"required,email"`
	Password   string     `json:"password,omitempty" binding:"required,min=8"`
	Role       string     `json:"role,omitempty" binding:"required"`
	CreatedAt  time.Time  `json:"created_at"`
	ArchivedAt *time.Time `json:"archived_at,omitempty"`
}
