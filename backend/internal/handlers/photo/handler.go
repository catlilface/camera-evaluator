package photo

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime/types"
	"net/http"
	photoApi "photo-upload-service/internal/pkg/api/photo"
	"photo-upload-service/internal/service/photo"
)

type photoService interface {
	ProcessPhoto(ctx context.Context, file types.File) (*photoApi.UploadResponse, error)
}

type UploadHandler struct {
	photoService photoService
}

func NewPhotoHandler(photoService *photo.Service) *UploadHandler {
	return &UploadHandler{
		photoService: photoService,
	}
}

func (h *UploadHandler) UploadPhoto(c *gin.Context) {
	ctx := c.Request.Context()
	var req photoApi.UploadPhotoMultipartBody

	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, photoApi.Error{
			Code:    "internal_server_error",
			Message: err.Error(),
		})
		return
	}

	res, err := h.photoService.ProcessPhoto(ctx, req.Photo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, photoApi.Error{
			Code:    "internal_server_error",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, res)
}
