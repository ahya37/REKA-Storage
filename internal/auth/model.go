package auth

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Email     string             `bson:"email"`
	Password  string             `bson:"password"`
	Role      string             `bson:"role"`
	IsActive  bool               `bson:"is_active"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}
