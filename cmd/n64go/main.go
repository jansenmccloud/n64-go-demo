package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"n64demo/assets"

	"github.com/clktmr/n64/drivers/controller"
	"github.com/clktmr/n64/drivers/display"
	"github.com/clktmr/n64/drivers/draw"
	"github.com/clktmr/n64/drivers/rspq/mixer"
	"github.com/clktmr/n64/fonts/gomono12"
	_ "github.com/clktmr/n64/machine"
	"github.com/clktmr/n64/rcp/audio"
	"github.com/clktmr/n64/rcp/serial/joybus"
	"github.com/clktmr/n64/rcp/texture"
	"github.com/clktmr/n64/rcp/video"
)

const (
	deInterlacing = true
)

var (
	// TODO load constants and variables from configuration
	font       = gomono12.NewFace()
	colorWheat = color.RGBA{245, 210, 174, 0}
	colorBlack = color.RGBA{1, 1, 1, 0}
)

// TODO after refactoring, think about creating an n64 go archetype setup with project skeleton
// TODO what about testability?
func main() {

	// TODO create consumer interface and extract code to impls => main should be as small as possible

	video.Setup(!deInterlacing)
	display := display.NewDisplay(image.Pt(320, 240), video.BPP16) // TODO experiment with 32bit and textture details

	// read inputs // TODO => abstraction needed
	controllers := make(chan [4]controller.Controller)
	go func() {
		var states [4]controller.Controller
		for {
			controller.Poll(&states)
			controllers <- states
		}
	}()

	// loading textures into memory // TODO => abstraction
	gopherFile, err := assets.Files.Open("anim/gopher-anim.CI8") // TODO build an abstraction so we don't need to know the filename here
	if err != nil {
		fmt.Printf("anim file error: %v", err)
		panic(err) // TODO think about proper error handling/ fallbacks
	}
	gopherTexture, err := texture.Load(gopherFile)
	if err != nil {
		fmt.Printf("texture error: %v", err)
		panic(err)
	}
	gopherRect := image.Rect(0, 0, 128, 128)
	blows := 0 // TODO something to inject into the game loop

	// play audio // TODO => abstraction
	audio.Start(48000)
	mixer.Init() // TODO experiment with different audio channels
	go func() {
		audio.Buffer.ReadFrom(mixer.Output)
	}()
	squeakFile, err := assets.Files.Open("sfx/squeak.pcm_s16be") // TODO abstraction to not need to know the file name here
	if err != nil {
		fmt.Printf("sfx file error: %v", err)
		panic(err)
	}
	squeakReader := squeakFile.(io.ReadSeeker)
	squeakSource := mixer.NewSource(squeakReader, 16000)

	// game loop
	for {
		fb := display.Swap()
		inputs := <-controllers // TODO think about a good way to inject instances into the draw loop

		// background
		draw.Src.Draw(fb, fb.Bounds(), &image.Uniform{colorWheat}, fb.Bounds().Min)

		// text
		text := fmt.Appendln(nil, "N64 - Demo")
		text = fmt.Appendf(text, "Buttons: %v\n", inputs[0].Down())
		text = fmt.Appendf(text, "Blows: %v/8\n\n", blows)
		textarea := fb.Bounds().Inset(15)
		pt := textarea.Min.Add(image.Pt(0, int(font.Ascent)+2))
		pt = draw.DrawText(fb, textarea, font, pt, &colorBlack, nil, text)

		// gopher anim
		gopherFrame := image.Point{} // sprite frame 0
		if blows < 8 {
			if inputs[0].Pressed()&joybus.ButtonA != 0 {
				squeakReader.Seek(0, io.SeekStart)
				mixer.SetSource(0, squeakSource)
				blows++
			}
			if inputs[0].Down()&joybus.ButtonA != 0 {
				gopherFrame.X += 128 // sprite frame 1
			}
		} else {
			gopherFrame.X += 256 // sprite frame 2 // TODO think about an behavior abstraction to cycle through the sprites => something for a lib
		}

		draw.Over.Draw(fb, gopherRect.Add(pt), gopherTexture, gopherFrame)

		draw.Flush()
	}
}
