package minio

import (
	"context"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/adityaeka26/deptech-test-backend/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Minio struct {
	minioClient *minio.Client
}

func NewMinio(config *config.EnvConfig) (*Minio, error) {
	minioClient, err := minio.New(config.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioAccessKey, config.MinioSecretKey, ""),
		Secure: config.MinioUseSSL,
	})
	if err != nil {
		return nil, err
	}

	return &Minio{
		minioClient: minioClient,
	}, nil
}

func (m *Minio) Upload(ctx context.Context, bucketName string, fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	objectName := fileHeader.Filename
	contentType := fileHeader.Header.Get("Content-Type")

	_, err = m.minioClient.PutObject(ctx, bucketName, objectName, file, fileHeader.Size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return err
	}

	return nil
}

func (m *Minio) GeneratePresignedURL(ctx context.Context, bucketName, path string, expiry time.Duration) (*url.URL, error) {
	return m.minioClient.PresignedGetObject(context.Background(), bucketName, path, expiry, make(url.Values))
}
