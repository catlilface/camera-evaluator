package photo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
	photoApi "photo-upload-service/internal/pkg/api/photo"
	"photo-upload-service/internal/rabbitmq"
)

type Service struct {
	queuePublisher *rabbitmq.Publisher
}

func NewPhotoService(queuePublisher *rabbitmq.Publisher) *Service {
	return &Service{
		queuePublisher: queuePublisher,
	}
}

func (s *Service) ProcessPhoto(ctx context.Context, file types.File) (*photoApi.UploadResponse, error) {
	id := uuid.New()

	bytes, err := file.Bytes()
	if err != nil {
		return nil, fmt.Errorf("service: failed read bytes: %w", err)
	}

	err = s.queuePublisher.PublishPhotoAsBase64(ctx, bytes, id)
	if err != nil {
		return nil, fmt.Errorf("service: failed to process file: %w", err)
	}

	return &photoApi.UploadResponse{ID: id}, nil
}
