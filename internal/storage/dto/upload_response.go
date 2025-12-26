package dto

import "time"

type UploadResponse struct {
	ID           string    `json:"id"`
	Filename     string    `json:"filename"`
	OriginalName string    `json:"original_name"`
	MimeType     string    `json:"mime_type"`
	Size         int64     `json:"size"`
	Bucket       string    `json:"bucket"`
	ObjectName   string    `json:"object_name"`
	URL          string    `json:"url"`
	CreatedAt    time.Time `json:"created_at"`
}
