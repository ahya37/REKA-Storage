package repositories

import (
	"context"
	"reka-storage/src/user/models"

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

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := r.db.Collection("users").FindOne(
		ctx,
		bson.M{"email": email},
	).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Insert(ctx context.Context, user *models.User) (*models.User, error) {

	res, err := r.db.Collection("users").InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.Id = res.InsertedID.(primitive.ObjectID)

	return user, nil
}
