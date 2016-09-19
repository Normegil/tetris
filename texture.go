package main

import "github.com/veandco/go-sdl2/sdl"

type texture struct {
	*sdl.Texture
	W, H int32
}

func (t *texture) Close() {
	t.Texture.Destroy()
}
