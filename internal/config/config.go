package config

import (
	"image/color"

	"github.com/clktmr/n64/fonts/gomono12"
)

// Config holds application configuration
type Config struct {
	Display DisplayConfig
	Colors  ColorConfig
}

// DisplayConfig holds display-related settings
type DisplayConfig struct {
	Width         int
	Height        int
	BitsPerPixel  int
	DeInterlacing bool
}

// ColorConfig holds color theme
type ColorConfig struct {
	Background color.RGBA
	Text       color.RGBA
}

// NewDefaultConfig returns default configuration
func NewDefaultConfig() *Config {
	return &Config{
		Display: DisplayConfig{
			Width:         320,
			Height:        240,
			BitsPerPixel:  16,
			DeInterlacing: true,
		},
		Colors: ColorConfig{
			Background: color.RGBA{245, 210, 174, 0},
			Text:       color.RGBA{1, 1, 1, 0},
		},
	}
}

// Font returns the default font face
func (c *Config) Font() interface{} {
	return gomono12.NewFace()
}
