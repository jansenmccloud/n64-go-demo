package main

import (
	"fmt"
	"image"

	"n64demo/internal/assets"
	"n64demo/internal/audio"
	"n64demo/internal/config"
	"n64demo/internal/game"
	"n64demo/internal/input"

	"github.com/clktmr/n64/drivers/display"
	"github.com/clktmr/n64/drivers/draw"
	_ "github.com/clktmr/n64/machine"
	"github.com/clktmr/n64/rcp/video"
)

func main() {
	// Initialize configuration
	cfg := config.NewDefaultConfig()
	
	// Setup video and display
	video.Setup(!cfg.Display.DeInterlacing)
	disp := display.NewDisplay(
		image.Pt(cfg.Display.Width, cfg.Display.Height),
		video.BPP16,
	)
	
	// Initialize managers
	inputMgr := input.NewManager()
	assetLoader := assets.NewLoader()
	audioMgr := audio.NewManager(assetLoader, 48000)
	
	// Load assets
	gopherTexture, err := assetLoader.LoadTexture("gopher-anim.CI8")
	if err != nil {
		fmt.Printf("Failed to load texture: %v\n", err)
		panic(err)
	}
	
	if err := audioMgr.LoadSound("squeak.pcm_s16be", 16000); err != nil {
		fmt.Printf("Failed to load sound: %v\n", err)
		panic(err)
	}
	
	// Initialize game state
	gameState := game.NewState()
	
	// Game loop
	for {
		fb := disp.Swap()
		inputs := inputMgr.GetInputs()
		
		// Update and render
		gameState.UpdateWithInput(inputs[0], audioMgr)
		gameState.Render(
			fb,
			cfg.Font(),
			&image.Uniform{cfg.Colors.Background},
			&cfg.Colors.Text,
			gopherTexture,
		)
		
		draw.Flush()
	}
}
