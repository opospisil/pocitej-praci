package config

import (
	"image/color"
)

// Config holds all configurable parameters for the game.
type Config struct {
	DisplayDurationMs int
	FontSize          int
	ColorPalette      []color.Color
}

// NewConfig returns a Config with sensible defaults.
func NewConfig() *Config {
	return &Config{
		DisplayDurationMs: 1500,
		FontSize:          300,
		ColorPalette: []color.Color{
			color.RGBA{R: 255, G: 0, B: 0, A: 255},   // Red
			color.RGBA{R: 0, G: 0, B: 255, A: 255},   // Blue
			color.RGBA{R: 0, G: 128, B: 0, A: 255},   // Green
			color.RGBA{R: 255, G: 255, B: 0, A: 255}, // Yellow
			color.RGBA{R: 128, G: 0, B: 128, A: 255}, // Purple
			color.RGBA{R: 255, G: 165, B: 0, A: 255}, // Orange
		},
	}
}
