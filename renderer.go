package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

type renderer struct {
	*sdl.Renderer
}

func (r *renderer) Text(font ttf.Font, text string, style sdl.Color) (*texture, error) {
	surface, err := font.RenderUTF8_Solid(text, style)
	if nil != err {
		return nil, err
	}
	defer surface.Free()

	t, err := r.CreateTextureFromSurface(surface)
	if nil != err {
		return nil, err
	}
	return &texture{t, surface.W, surface.H}, err
}

func (r *renderer) SetDrawColor(color sdl.Color) error {
	return r.Renderer.SetDrawColor(color.R, color.G, color.B, color.A)
}

func (r *renderer) Close() {
	r.Renderer.Destroy()
}
