package storage

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type File struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID       string             `bson:"user_id" json:"user_id"`
	Bucket       string             `bson:"bucket" json:"bucket"`
	Path         string             `bson:"path" json:"path"`
	Filename     string             `bson:"filename" json:"filename"`
	OriginalName string             `bson:"original_name" json:"original_name"`
	Size         int64              `bson:"size" json:"size"`
	Url          string             `bson:"url" json:"url"`
	MimeType     string             `bson:"mime_type" json:"mime_type"`
	CreatedAt    time.Time          `bson:"created_at" json:"uploaded_at"`
}
