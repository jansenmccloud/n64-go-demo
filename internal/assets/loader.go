package assets

import (
	"fmt"
	"n64demo/assets"

	"github.com/clktmr/n64/rcp/texture"
)

// Loader handles loading assets from embedded files
type Loader struct{}

// NewLoader creates a new asset loader
func NewLoader() *Loader {
	return &Loader{}
}

// LoadTexture loads a texture from the assets
func (al *Loader) LoadTexture(name string) (*texture.Image, error) {
	path := fmt.Sprintf("anim/%s", name)
	file, err := assets.Files.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open texture %s: %w", name, err)
	}
	
	tex, err := texture.Load(file)
	if err != nil {
		return nil, fmt.Errorf("failed to load texture %s: %w", name, err)
	}
	
	return tex, nil
}

// LoadSound opens a sound file from the assets
func (al *Loader) LoadSound(name string) (interface{}, error) {
	path := fmt.Sprintf("sfx/%s", name)
	file, err := assets.Files.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open sound %s: %w", name, err)
	}
	
	return file, nil
}
