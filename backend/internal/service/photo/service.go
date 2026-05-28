package photo

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"photo-upload-service/internal/models"
	photoApi "photo-upload-service/internal/pkg/api/photo"
	"photo-upload-service/internal/rabbitmq/producer"
	"strings"

	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
)

type Service struct {
	queuePublisher *producer.Publisher
	photosDir      string
}

func NewPhotoService(queuePublisher *producer.Publisher, photosDir string) *Service {
	return &Service{
		queuePublisher: queuePublisher,
		photosDir:      photosDir,
	}
}

func (s *Service) ProcessPhoto(ctx context.Context, data models.ProcessPhotoData) (*photoApi.EvaluateSuccessResponse, error) {
	var id uuid.UUID
	if data.ChannelID != "" {
		parsed, err := uuid.Parse(data.ChannelID)
		if err != nil {
			return nil, fmt.Errorf("service: invalid channel_id: %w", err)
		}
		id = parsed
	} else {
		id = uuid.New()
	}

	photosDir, err := s.resolvePhotosDir()
	if err != nil {
		return nil, err
	}

	ext := strings.ToLower(filepath.Ext(data.File.Filename()))
	if ext == "" {
		ext = ".png"
	}

	filePath := filepath.Join(photosDir, id.String()+ext)

	if err := os.MkdirAll(photosDir, 0755); err != nil {
		return nil, fmt.Errorf("service: failed to create directory: %w", err)
	}

	err = saveFile(filePath, data.File)
	if err != nil {
		return nil, fmt.Errorf("service: failed to save file: %w", err)
	}

	err = s.queuePublisher.PublishPhoto(ctx, id.String(), filePath)
	if err != nil {
		return nil, fmt.Errorf("service: failed to process file: %w", err)
	}

	return &photoApi.EvaluateSuccessResponse{ID: &id}, nil
}

func (s *Service) resolvePhotosDir() (string, error) {
	if s.photosDir != "" {
		return s.photosDir, nil
	}

	projectRoot, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("service: failed to get working directory: %w", err)
	}

	return filepath.Join(projectRoot, "..", ".photos"), nil
}

func saveFile(filePath string, file types.File) error {
	data, err := file.Bytes()
	if err != nil {
		return fmt.Errorf("error save file: %w", err)
	}
	return os.WriteFile(filePath, data, 0644)
}
