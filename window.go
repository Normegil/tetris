package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl"
)

type window struct {
	*sdl.Window
	render *Renderer
}

func newWindow(title string, def sdl.Rect, winFlags uint32, rendererFlags uint32) (*window, error) {
	logrus.WithFields(logrus.Fields{
		"Title":         title,
		"Position/Size": def,
	}).Info("Creating Window")
	sdlWindow, err := sdl.CreateWindow(title, int(def.X), int(def.Y), int(def.W), int(def.H), winFlags)
	if nil != err {
		return nil, err
	}
	window := &window{Window: sdlWindow, render: nil}

	render, err := sdl.CreateRenderer(window.Window, -1, rendererFlags)
	if nil != err {
		return nil, err
	}
	window.render = NewRenderer(render)
	window.render.SetDrawColor(sdl.Color{R: 0, G: 0, B: 0, A: 255})

	return window, nil
}

func (w *window) GetSize() Size {
	width, height := w.Window.GetSize()
	return Size{
		W: int32(width),
		H: int32(height),
	}
}

func (w *window) Renderer() *Renderer {
	return w.render
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
