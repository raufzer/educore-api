package response

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserResponse struct {
	ID        primitive.ObjectID    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type LoginResponse struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}
