# Development

- [Development](#development)
  - [Initial setup environment](#initial-setup-environment)
  - [Build and run n64go project](#build-and-run-n64go-project)
  - [Asset management](#asset-management)
    - [Textures - PNG to CI8](#textures---png-to-ci8)
    - [Sound FX - ss16be](#sound-fx---ss16be)

## Initial setup environment

* have a go dev environment
* have ares emulator installed and its executable in your path
* `go install github.com/embeddedgo/dl/go1.24.4-embedded@latest`
* `go1.24.4-embedded download`
* `go install github.com/clktmr/n64/tools/n64go@v0.1.2`
* for code completion in your IDE start gopls in the embedded-go environment (GOENV=go.env)
  * e.g in VScode add settings: `"gopls": {"build.env": {"GOENV":"go.env"}}`  
  
## Build and run n64go project

> **Info:** Building and running your project requires to have the embedded-go toolchain in your GOENV

To clean up, to build and to run the n64go project execute the following from project root directory: 
* `./buildrun.sh`

## Asset management

### Textures - PNG to CI8

The n64 texture format CI8 stores a palette of up to 256 colors. To convert from PNG use the the following command:

`n64go texture -format CI8 -palette 128 ./assets/raw/gopher-anim.png`

### Sound FX - ss16be

Supported sound format is a `.pcm_s16be` file with uncompressed mono 16KHz sample rate. Use `ffmpeg` for conversion:

`ffmpeg -i <infile> -ac 1 -ar <samplerate> -f s16be -c:a pcm_s16be <outfile>`