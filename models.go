package main

import (
	"github.com/google/uuid"
	"github.com/oguzcoll/rssagg/internal/database"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func databaseUserToUser(u database.User) User {
	return User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Name:      u.Name,
	}
}
