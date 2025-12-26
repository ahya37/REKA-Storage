package storage

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/minio/minio-go/v7"
)

type Service struct {
	minio  *minio.Client
	bucket string
	repo   *Repository
}

func NewService(minio *minio.Client, bucket string, repo *Repository) *Service {
	return &Service{
		minio:  minio,
		bucket: bucket,
		repo:   repo,
	}
}

func (s *Service) Upload(
	ctx context.Context,
	userID string,
	file *multipart.FileHeader,
	folder string,
) (*File, error) {

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	ext := filepath.Ext(file.Filename)

	randomName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)

	objectPath := folder + "/" + randomName

	_, err = s.minio.PutObject(
		ctx,
		s.bucket,
		objectPath,
		src,
		file.Size,
		minio.PutObjectOptions{
			ContentType: file.Header.Get("Content-Type"),
		},
	)
	if err != nil {
		return nil, err
	}

	fileModel := &File{
		UserID:       userID,
		Bucket:       s.bucket,
		Path:         objectPath,
		Filename:     randomName,
		OriginalName: file.Filename,
		Size:         file.Size,
		Url:          "https://" + s.minio.EndpointURL().Host + "/" + objectPath,
		MimeType:     file.Header.Get("Content-Type"),
		CreatedAt:    time.Now(),
	}

	if err := s.repo.Insert(ctx, fileModel); err != nil {
		return nil, err
	}

	return fileModel, nil
}
