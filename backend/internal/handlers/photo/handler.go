package photo

import (
	"context"
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
}

func NewPhotoHandler(photoService *photo.Service) *UploadHandler {
	return &UploadHandler{
		photoService: photoService,
	}
}

func (h *UploadHandler) Evaluate(c *gin.Context) {
	method := c.PostForm("method_id")

	fileHeader, err := c.FormFile("image")
	if err != nil {
		httpUtils.AbortWithStatus(c, http.StatusBadRequest, err)
		return
	}

	var file types.File
	file.InitFromMultipart(fileHeader)

	res, err := h.photoService.ProcessPhoto(c.Request.Context(), models.ProcessPhotoData{
		File:   file,
		Method: method,
	})
	if err != nil {
		httpUtils.AbortWithStatus(c, http.StatusInternalServerError, err)
		return
	}

	httpUtils.Success(c, res)
}
