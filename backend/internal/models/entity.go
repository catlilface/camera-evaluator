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
}

type ProcessPhotoData struct {
	File      types.File
	Method    string
	DisplayID int
}
