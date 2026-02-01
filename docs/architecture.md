# N64 Go Demo - Package Structure

## Project Structure

```
n64-go-demo/
├── cmd/
│   └── n64go/
│       └── main.go          # Application entry point
├── internal/
│   ├── assets/
│   │   └── loader.go        # Asset loading (textures, sounds)
│   ├── audio/
│   │   └── manager.go       # Audio playback and mixing
│   ├── config/
│   │   └── config.go        # Application configuration
│   ├── game/
│   │   └── state.go         # Game state and rendering logic
│   └── input/
│       └── controller.go    # Controller input management
├── assets/
│   ├── anim/                # Animation textures
│   ├── sfx/                 # Sound effects
│   └── assets.go            # Embedded assets
└── go.mod
```

## Package Overview

### `cmd/n64go`
Main application entry point. Keeps initialization logic minimal and delegates to internal packages.

### `internal/config`
Configuration management:
- Display settings (resolution, color depth)
- Color themes
- Font configuration

### `internal/input`
Controller input abstraction:
- Polling N64 controllers
- Input state management
- Non-blocking input handling

### `internal/assets`
Asset loading:
- Texture loading from embedded files
- Sound file loading
- Abstraction over file paths

### `internal/audio`
Audio system:
- Audio manager for playback
- Sound source management
- Multi-channel audio mixing

### `internal/game`
Game logic:
- Game state management
- Input handling
- Rendering logic
- Sprite animation

## Design Principles

1. **Separation of Concerns**: Each package has a single, well-defined responsibility
2. **Dependency Injection**: Dependencies are passed via interfaces where appropriate
3. **Testability**: Business logic is decoupled from N64-specific APIs
4. **Go Best Practices**: Following standard Go project layout with `internal/` packages

## Usage

The main application initializes all managers and runs the game loop:

```go
cfg := config.NewDefaultConfig()
inputMgr := input.NewManager()
assetLoader := assets.NewLoader()
audioMgr := audio.NewManager(assetLoader, 48000)
gameState := game.NewState()

// Game loop
for {
    inputs := inputMgr.GetInputs()
    gameState.UpdateWithInput(inputs[0], audioMgr)
    gameState.Render(fb, cfg, gopherTexture)
}
```
