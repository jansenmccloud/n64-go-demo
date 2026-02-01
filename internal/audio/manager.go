package audio

import (
	"fmt"
	"io"

	"github.com/clktmr/n64/drivers/rspq/mixer"
	"github.com/clktmr/n64/rcp/audio"
)

// AssetLoader interface for loading sound files
type AssetLoader interface {
	LoadSound(name string) (interface{}, error)
}

// Manager handles audio playback and mixing
type Manager struct {
	assetLoader AssetLoader
	sources     map[string]*Source
}

// Source represents a loaded audio source
type Source struct {
	reader io.ReadSeeker
	source *mixer.Source
}

// NewManager creates a new audio manager
func NewManager(assetLoader AssetLoader, sampleRate int) *Manager {
	audio.Start(sampleRate)
	mixer.Init()
	
	go func() {
		audio.Buffer.ReadFrom(mixer.Output)
	}()
	
	return &Manager{
		assetLoader: assetLoader,
		sources:     make(map[string]*Source),
	}
}

// LoadSound loads a sound effect by name
func (am *Manager) LoadSound(name string, sampleRate int) error {
	file, err := am.assetLoader.LoadSound(name)
	if err != nil {
		return fmt.Errorf("failed to load sound %s: %w", name, err)
	}
	
	reader, ok := file.(io.ReadSeeker)
	if !ok {
		return fmt.Errorf("sound file %s does not support seeking", name)
	}
	
	source := mixer.NewSource(reader, sampleRate)
	am.sources[name] = &Source{
		reader: reader,
		source: source,
	}
	
	return nil
}

// PlaySound plays a sound on the specified channel
func (am *Manager) PlaySound(name string, channel int) error {
	src, exists := am.sources[name]
	if !exists {
		return fmt.Errorf("sound %s not loaded", name)
	}
	
	src.reader.Seek(0, io.SeekStart)
	mixer.SetSource(channel, src.source)
	
	return nil
}
