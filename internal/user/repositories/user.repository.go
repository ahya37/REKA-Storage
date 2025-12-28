package repositories

import (
	"context"
	"reka-storage/internal/user/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Database
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{db: db}

}

func (r *UserRepository) FindByID(ctx context.Context, userID string) (*models.User, error) {
	var user models.User

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	err = r.db.Collection("users").FindOne(
		ctx,
		bson.M{"_id": id},
	).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
