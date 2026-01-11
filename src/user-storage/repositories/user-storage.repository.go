package repositories

import (
	"context"
	"reka-storage/src/user-storage/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStorageRepository struct {
	db *mongo.Database
}

func NewUserStorageRepository(db *mongo.Database) *UserStorageRepository {
	return &UserStorageRepository{db: db}
}

func (r *UserStorageRepository) GetUsage(
	ctx context.Context,
	userID string,
) (*models.UserStorage, error) {
	var usage models.UserStorage

	err := r.db.Collection("user_storages").FindOne(
		ctx,
		bson.M{"user_id": userID},
	).Decode(&usage)

	if err != nil {
		return nil, err
	}

	return &usage, nil

}
