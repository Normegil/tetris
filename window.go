package main

import "github.com/veandco/go-sdl2/sdl"

type window struct {
	*sdl.Window
	render *renderer
}

func newWindow(title string, def sdl.Rect, flags uint32) (*window, error) {
	win, err := sdl.CreateWindow(title, int(def.X), int(def.Y), int(def.W), int(def.H), flags)
	if nil != err {
		return nil, err
	}

	return &window{Window:win, render:nil}, nil
}

func (w *window) GetSize() Size {
	width, height := w.Window.GetSize()
	return Size{
		W: int32(width),
		H: int32(height),
	}
}

func (w *window) Renderer() (*renderer, error) {
	if nil == w.render {
		render, err := sdl.CreateRenderer(w.Window, -1, sdl.RENDERER_ACCELERATED)
		if nil != err {
			return nil, err
		}
		w.render = &renderer{render}
		w.render.SetDrawColor(sdl.Color{R: 0, G: 0, B: 0, A: 255})
	}
	return w.render, nil
}

func (w *window) Close() {
	if nil != w.render {
		w.render.Close()
	}
	w.Window.Destroy()
}

type Size struct {
	W, H int32
}