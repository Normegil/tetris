package sdl

import (
	"github.com/normegil/log"
	"github.com/veandco/go-sdl2/sdl"
)

// Window decorate an sdl.Window, containing renderer and easier/more consistant method to use. Use NewWindow() to build
type Window struct {
	*sdl.Window
	render *Renderer

	Logger log.AgnosticLogger
}

// NewWindow is the Window constructor, using winFlags flags. It will create the associated renderer (that can be obtain via Renderer()), using rendererFlags
func NewWindow(title string, def sdl.Rect, winFlags uint32, rendererFlags uint32) (*Window, error) {
	sdlWindow, err := sdl.CreateWindow(title, int(def.X), int(def.Y), int(def.W), int(def.H), winFlags)
	if nil != err {
		return nil, err
	}
	window := &Window{Window: sdlWindow, render: nil}

	render, err := sdl.CreateRenderer(window.Window, -1, rendererFlags)
	if nil != err {
		return nil, err
	}
	window.render = NewRenderer(render)
	err = window.render.SetDrawColor(sdl.Color{R: 0, G: 0, B: 0, A: 255})
	if nil != err {
		return nil, err
	}

	return window, nil
}

// Size will send back the size of the current window
func (w *Window) Size() Size {
	width, height := w.Window.GetSize()
	return Size{
		W: int32(width),
		H: int32(height),
	}
}

// Renderer returns the Window associated renderer
func (w *Window) Renderer() *Renderer {
	return w.render
}

// Close the current window
func (w *Window) Close() {
	if nil != w.render {
		w.render.Close()
	}
	w.Window.Destroy()
}

// Size is a simple struct for 2D dimensions
type Size struct {
	W, H int32
}
