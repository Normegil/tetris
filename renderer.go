package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type TextStyle struct {
	FontName string
	FontSize int
	Color    sdl.Color
}

type TextStyleWithPos struct {
	TextStyle
	Position sdl.Point
}

type Renderer struct {
	*sdl.Renderer
	fonts *fonts
}

func NewRenderer(renderer *sdl.Renderer) *Renderer {
	return &Renderer{
		Renderer: renderer,
		fonts:    &fonts{},
	}
}

func (r *Renderer) Text(text string, style TextStyleWithPos) error {
	font, err := r.fonts.load(style.FontName).size(style.FontSize)
	if nil != err {
		return err
	}

	surface, err := font.RenderUTF8_Solid(text, style.Color)
	if nil != err {
		return err
	}
	defer surface.Free()

	t, err := r.CreateTextureFromSurface(surface)
	if nil != err {
		return err
	}
	defer t.Destroy()

	return r.Copy(t, &sdl.Rect{
		W: surface.W,
		H: surface.H,
	}, &sdl.Rect{
		X: style.Position.X,
		Y: style.Position.Y,
		W: surface.W,
		H: surface.H,
	})
}

func (r *Renderer) TextureSize(text string, style TextStyle) (Size, error) {
	font, err := r.fonts.load(style.FontName).size(style.FontSize)
	if nil != err {
		return Size{}, err
	}

	surface, err := font.RenderUTF8_Solid(text, style.Color)
	if nil != err {
		return Size{}, err
	}
	defer surface.Free()

	return Size{
		W: surface.W,
		H: surface.H,
	}, nil
}

func (r *Renderer) SetDrawColor(color sdl.Color) error {
	return r.Renderer.SetDrawColor(color.R, color.G, color.B, color.A)
}

func (r *Renderer) Close() {
	r.Renderer.Destroy()
}
