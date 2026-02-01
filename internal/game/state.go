package game

import (
	"fmt"
	"image"
	"image/draw"

	"github.com/clktmr/n64/drivers/controller"
	drawpkg "github.com/clktmr/n64/drivers/draw"
	"github.com/clktmr/n64/rcp/serial/joybus"
	"github.com/clktmr/n64/rcp/texture"
)

// Config interface for accessing configuration
type Config interface {
	Font() interface{}
	GetColors() (background, text image.Image)
}

// AudioManager interface for audio playback
type AudioManager interface {
	PlaySound(name string, channel int) error
}

// State holds the current game state
type State struct {
	gopherRect   image.Rectangle
	blows        int
	maxBlows     int
	currentInput controller.Controller
}

// NewState creates a new game state
func NewState() *State {
	return &State{
		gopherRect: image.Rect(0, 0, 128, 128),
		blows:      0,
		maxBlows:   8,
	}
}

// UpdateWithInput updates the game state and handles input
func (g *State) UpdateWithInput(input controller.Controller, audioMgr AudioManager) {
	g.currentInput = input
	
	if g.blows < g.maxBlows {
		if input.Pressed()&joybus.ButtonA != 0 {
			audioMgr.PlaySound("squeak.pcm_s16be", 0)
			g.blows++
		}
	}
}

// Render renders the current game state
func (g *State) Render(fb draw.Image, font interface{}, bgColor, textColor image.Image, gopherTexture *texture.Image) {
	// Background
	drawpkg.Src.Draw(fb, fb.Bounds(), bgColor, fb.Bounds().Min)
	
	// Text
	text := fmt.Appendln(nil, "N64 - Demo")
	text = fmt.Appendf(text, "Buttons: %v\n", g.currentInput.Down())
	text = fmt.Appendf(text, "Blows: %v/%v\n\n", g.blows, g.maxBlows)
	
	textarea := fb.Bounds().Inset(15)
	pt := textarea.Min.Add(image.Pt(0, 12+2)) // font ascent approximation
	pt = drawpkg.DrawText(fb, textarea, font, pt, textColor, nil, text)
	
	// Gopher animation
	gopherFrame := g.getGopherFrame()
	drawpkg.Over.Draw(fb, g.gopherRect.Add(pt), gopherTexture, gopherFrame)
}

// getGopherFrame returns the current sprite frame offset
func (g *State) getGopherFrame() image.Point {
	if g.blows >= g.maxBlows {
		return image.Pt(256, 0) // frame 2 - finished
	}
	
	if g.currentInput.Down()&joybus.ButtonA != 0 {
		return image.Pt(128, 0) // frame 1 - pressed
	}
	
	return image.Pt(0, 0) // frame 0 - idle
}
