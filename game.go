package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

type game struct{}

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

	f := fonts{}
	defer f.Close()

	win, err := newWindow("Tetris", sdl.Rect{}, sdl.WINDOW_SHOWN)
	if nil != err {
		return err
	}
	defer win.Destroy()

	var screen screen
	screen = MenuScreen{}
	ctrl := loopCtrl{quit: false, fps: fpsControls{capped: true, number: 60}}
	return loop(ctrl, func(loop loopCtrl) (loopCtrl, error) {
		screen, err = screen.execute(win, &f)
		if nil != err {
			return loopCtrl{quit: true}, err
		}
		return loopCtrl{quit:screen.quit(), fps:loop.fps}, nil
	})
}
