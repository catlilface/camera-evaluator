package photo

import (
	"context"
	"fmt"
	"net/http"
	"photo-upload-service/internal/models"
	api "photo-upload-service/internal/pkg/api/photo"
	"photo-upload-service/internal/service/photo"
	httpUtils "photo-upload-service/pkg/utils/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime/types"
)

type photoService interface {
	ProcessPhoto(ctx context.Context, data models.ProcessPhotoData) (*api.EvaluateSuccessResponse, error)
}

type UploadHandler struct {
	photoService photoService
	maxFileSize  int64
}

func NewPhotoHandler(photoService *photo.Service, maxFileSize int64) *UploadHandler {
	return &UploadHandler{
		photoService: photoService,
		maxFileSize:  maxFileSize,
	}
}

func (h *UploadHandler) Evaluate(c *gin.Context) {
	method := c.PostForm("method_id")
	channelID := c.PostForm("channel_id")

	fileHeader, err := c.FormFile("image")
	if err != nil {
		httpUtils.AbortWithStatus(c, http.StatusBadRequest, err)
		return
	}

	if fileHeader.Size > h.maxFileSize {
		httpUtils.AbortWithStatus(c, http.StatusRequestEntityTooLarge, fmt.Errorf("file size exceeds maximum allowed size of %d bytes", h.maxFileSize))
		return
	}

	var file types.File
	file.InitFromMultipart(fileHeader)

	res, err := h.photoService.ProcessPhoto(c.Request.Context(), models.ProcessPhotoData{
		File:      file,
		Method:    method,
		ChannelID: channelID,
	})
	if err != nil {
		httpUtils.AbortWithStatus(c, http.StatusInternalServerError, err)
		return
	}

	httpUtils.Success(c, res)
}
