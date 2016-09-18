package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

type game struct {}

func newGame() *game {
	return &game{}
}

func (g *game) run() error {
	err := sdl.Init(sdl.INIT_VIDEO)
	if nil != err {
		return err
	}
	defer sdl.Quit()

	err = ttf.Init()
	if nil != err {
		return err
	}
	defer ttf.Quit()

	win, err := newWindow("Tetris", sdl.Rect{}, sdl.WINDOW_SHOWN)
	if nil != err {
		return err
	}
	defer win.Destroy()

	screen := MENU
	return loop(loopCtrl{fpsControls{capped:true, number: 60}}, func(loop loopCtrl) (loopCtrl, error) {
		screen, err = screen()
		if nil != err {
			return nil, err
		}
		if NONE == screen {
			return loopCtrl{quit:true}
		}
		return loop
	})
}