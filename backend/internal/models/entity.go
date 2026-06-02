package models

import (
	"github.com/oapi-codegen/runtime/types"
)

type EvaluationMessage struct {
	ID        string `json:"id"`
	ImagePath string `json:"image_path"`
}

type ProcessPhotoData struct {
	File      types.File
	Method    string
	ChannelID string
}

type EvaluationResponseMessage struct {
	ID            string  `json:"id"`
	Score         float64 `json:"score"`
	AttentionPath string  `json:"attn_img"`
}
