package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson:"_id"`

	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	Password  string    `bson:"password"`
	Role      string    `bson:"role"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

// // models/user.go (if using GORM)
// package models

// import (
//   "time"

//   "gorm.io/gorm"
// )

// type User struct {
//   gorm.Model
//   Username string `gorm:"uniqueIndex;not null"` // Unique username for identification
//   // OR
//   Email    string `gorm:"uniqueIndex;not null"` // Unique email for identification
//   Password string `gorm:"not null"`
//   Role     string `json:"role,omitempty"` // Optional field for roles
//   CreatedAt time.Time
//   UpdatedAt time.Time
// }
