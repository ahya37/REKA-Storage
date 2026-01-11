package models

import "time"

type UserStorage struct {
	ID         string    `bson:"_id,omitempty"`
	UserID     string    `bson:"user_id"`
	UsedBytes  int64     `bson:"used_bytes"`
	QuotaBytes int64     `bson:"quota_bytes"`
	UpdatedAt  time.Time `bson:"updated_at"`
}
