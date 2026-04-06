package photo

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
	"os"
	"path/filepath"
	"photo-upload-service/internal/models"
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

func (s *Service) ProcessPhoto(ctx context.Context, data models.ProcessPhotoData) (*photoApi.EvaluateSuccessResponse, error) {
	id := uuid.New()

	projectRoot, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("service: failed to get working directory: %w", err)
	}

	photosDir := filepath.Join(projectRoot, "..", "photosInProcess") //TODO: путь из конфига?
	filePath := filepath.Join(photosDir, id.String()+".png")

	if err := os.MkdirAll(photosDir, 0755); err != nil {
		return nil, fmt.Errorf("service: failed to create directory: %w", err)
	}

	err = saveFile(filePath, data.File)
	if err != nil {
		return nil, fmt.Errorf("service: failed to save file: %w", err)
	}

	err = s.queuePublisher.PublishPhoto(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("service: failed to process file: %w", err)
	}

	return &photoApi.EvaluateSuccessResponse{ID: &id}, nil
}

func saveFile(filePath string, file types.File) error {
	data, err := file.Bytes()
	if err != nil {
		return fmt.Errorf("error save file: %w", err)
	}
	return os.WriteFile(filePath, data, 0644)
}
