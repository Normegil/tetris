package main

import "github.com/veandco/go-sdl2/sdl"

type window struct {
	*sdl.Window
}

func newWindow(title string, def sdl.Rect, flags uint32) (*window, error) {
	window, err := sdl.CreateWindow(title, def.X, def.Y, def.W, def.H, flags)
	if nil != err {
		return nil, err
	}
	return window{window}, nil
}