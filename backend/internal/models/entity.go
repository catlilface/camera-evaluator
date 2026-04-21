package models

import (
	"github.com/google/uuid"
	"github.com/oapi-codegen/runtime/types"
)

type MonitorData struct {
	ID           int     `json:"ID"`
	PixelDensity int     `json:"pixel-density"`
	Brightness   float64 `json:"brightness"`
	Contrast     float64 `json:"contrast"`
	DeltaE       float64 `json:"delta_e"`
	ColorDepth   float64 `json:"color_depth"`
	Temperature  int     `json:"temperature"`
	Gamma        float64 `json:"gamma"`
	Default      bool    `json:"default"`
}

type EvaluationMessage struct {
	PhotoID uuid.UUID   `json:"photo_id"`
	Monitor MonitorData `json:"monitor"`
	Method  string      `json:"method"`
	Ext     string      `json:"ext"`
}

type ProcessPhotoData struct {
	File   types.File
	Method string
}

type EvaluationResponseMessage struct {
	ImageID string      `json:"image_id"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}
