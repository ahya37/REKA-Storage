package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Database
}

func NewRepository(collection *mongo.Database) *Repository {
	return &Repository{collection: collection}
}

func (r *Repository) Insert(
	ctx context.Context,
	file *File,
) error {
	res, err := r.collection.Collection("files").InsertOne(ctx, file)
	if err != nil {
		return err
	}

	file.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *Repository) ListByUser(ctx context.Context, userID string) ([]*File, error) {
	filter := bson.M{
		"user_id": userID,
	}

	cursor, err := r.collection.Collection("files").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var files []*File
	if err := cursor.All(ctx, &files); err != nil {
		return nil, err
	}

	return files, nil
}
