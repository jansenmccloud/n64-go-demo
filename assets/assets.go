package assets

import (
	"embed"

	"github.com/clktmr/n64/drivers/cartfs"
)

var (
	//go:embed anim/gopher-anim.CI8 sfx/squeak.pcm_s16be
	files embed.FS
  Files cartfs.FS = cartfs.Embed(files) // TODO think about when to load resources
)
