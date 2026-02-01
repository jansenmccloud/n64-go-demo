# GitHub Copilot Instructions for N64 Go Demo

## Project Overview

This is a Nintendo 64 ROM project written in Go using the [n64](https://github.com/clktmr/n64) library. The project demonstrates basic N64 game development including graphics rendering, controller input, audio playback, and sprite animations.

## Project Structure

```
n64-go-demo/
├── cmd/n64go/
│   └── main.go              # Application entry point - minimal initialization only
├── internal/                # Private packages (Go best practice)
│   ├── config/              # Configuration management
│   ├── input/               # Controller input handling
│   ├── assets/              # Asset loading (textures, sounds)
│   ├── audio/               # Audio playback and mixing
│   └── game/                # Game state and rendering logic
├── assets/                  # Embedded assets (textures, sounds)
└── docs/                    # Documentation
```

## Architecture Principles

### 1. Separation of Concerns
- **Each package has ONE responsibility**
- `cmd/n64go/main.go` should remain minimal - only initialization and game loop
- Business logic lives in `internal/` packages

### 2. Dependency Injection
- Pass dependencies via constructor functions (e.g., `NewManager(deps...)`)
- Use interfaces for testability (e.g., `audio.AssetLoader`)
- Avoid global variables except for truly global state

### 3. Package Organization
- **internal/config**: All configuration structs and defaults
- **internal/input**: Controller polling and input state management
- **internal/assets**: Loading textures and sounds from embedded files
- **internal/audio**: Audio manager, sound sources, mixing
- **internal/game**: Game state, update logic, rendering

## Coding Conventions

### Package Names
- Use short, lowercase, single-word names
- No underscores or camelCase
- Package name should match directory name

### Type Naming
- Export types that are part of the public API: `config.Config`, `input.Manager`
- Use descriptive names: `State` not `S`, `Manager` not `Mgr` (except in variables)
- Constructor pattern: `NewXxx()` returns `*Xxx`

### Variable Naming
- Short names in small scopes: `cfg`, `fb`, `pt`
- Longer names for package-level or exported items
- Receivers: 1-2 letter abbreviations (e.g., `g *State`, `am *Manager`)

### Error Handling
- Always check errors from N64 library functions
- Use `fmt.Errorf` with `%w` for error wrapping
- Log errors before panicking (for development)
- Consider graceful degradation for production

## N64-Specific Guidelines

### Display & Video
- Standard resolution: 320x240 (configurable)
- Use `video.BPP16` for 16-bit color depth
- Double buffering: call `display.Swap()` each frame
- Always call `draw.Flush()` after rendering operations

### Controller Input
- Poll controllers in a goroutine (see `internal/input/controller.go`)
- Use channels for thread-safe state sharing
- Button states: `Pressed()` for new presses, `Down()` for held buttons
- Button constants from `github.com/clktmr/n64/rcp/serial/joybus`

### Textures & Sprites
- Load textures via `texture.Load()` from embedded files
- Sprite sheets use `image.Point` for frame offsets
- Frame size is fixed per sprite (e.g., 128x128 for gopher)
- Use `draw.Over` for transparent sprites, `draw.Src` for backgrounds

### Audio
- Initialize once: `audio.Start(sampleRate)` and `mixer.Init()`
- Run mixer in goroutine: `audio.Buffer.ReadFrom(mixer.Output)`
- Load sounds as `io.ReadSeeker` for repeatable playback
- Use `mixer.SetSource(channel, source)` to play sounds
- Seek to 0 before replaying: `reader.Seek(0, io.SeekStart)`

## Common Patterns

### Initialization in main.go
```go
cfg := config.NewDefaultConfig()
inputMgr := input.NewManager()
assetLoader := assets.NewLoader()
audioMgr := audio.NewManager(assetLoader, 48000)
gameState := game.NewState()
```

### Game Loop Structure
```go
for {
    fb := display.Swap()           // Get framebuffer
    inputs := inputMgr.GetInputs() // Get controller state
    
    gameState.UpdateWithInput(inputs[0], audioMgr)
    gameState.Render(fb, cfg, textures...)
    
    draw.Flush()                   // Commit drawing operations
}
```

### Loading Assets
```go
// Textures
texture, err := assetLoader.LoadTexture("filename.CI8")

// Sounds
err := audioMgr.LoadSound("filename.pcm_s16be", sampleRate)
```

## Testing Considerations

- N64-specific code cannot be unit tested directly (requires hardware/emulator)
- Extract testable logic into pure functions where possible
- Use interfaces to mock N64 dependencies (display, audio, input)
- Focus integration tests on game logic in `internal/game`

## Performance Notes

- N64 is resource-constrained (4MB RAM, 93.75 MHz CPU)
- Minimize allocations in the game loop
- Reuse buffers and avoid `append()` in hot paths
- Profile with emulators before optimizing

## When Adding New Features

1. **Determine the right package**: Does it fit existing packages or need a new one?
2. **Define interfaces first**: What does this component need from others?
3. **Keep main.go clean**: Initialization only, no business logic
4. **Update architecture.md**: Document new packages and patterns
5. **Consider N64 constraints**: Memory, CPU, and API limitations

## Resources

- [N64 Go Library](https://github.com/clktmr/n64)
- [Tutorial Series](https://www.timurcelik.de/posts/n64go-1-getting-started/)
- [Project Architecture](../docs/architecture.md)

## Common Mistakes to Avoid

❌ Don't put business logic in `main.go`
❌ Don't use global state (except for truly global resources)
❌ Don't forget `draw.Flush()` after rendering
❌ Don't allocate in game loop hot paths
❌ Don't ignore errors from N64 library calls
❌ Don't use blocking operations in the game loop

✅ Do use dependency injection
✅ Do keep packages focused and cohesive
✅ Do handle errors appropriately
✅ Do document N64-specific quirks
✅ Do follow Go best practices and idioms
