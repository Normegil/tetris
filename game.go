package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/normegil/sdl"
	"github.com/normegil/sdl/games"
	sdl2 "github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

type game struct{}

func newGame() *game {
	return &game{}
}

func (g *game) run() error {
	const FPS = 60

	err := sdl2.Init(sdl2.INIT_VIDEO)
	if nil != err {
		return err
	}
	defer logrus.Debug("SDL(with TTF) Exited")
	defer sdl2.Quit()

	err = ttf.Init()
	if nil != err {
		return err
	}
	defer ttf.Quit()
	logrus.Debug("SDL(with TTF) Launched")

	win, err := sdl.NewWindow("Tetris", sdl2.Rect{
		X: sdl2.WINDOWPOS_UNDEFINED,
		Y: sdl2.WINDOWPOS_UNDEFINED,
		W: 640,
		H: 480,
	}, sdl2.WINDOW_SHOWN|sdl2.WINDOW_FULLSCREEN_DESKTOP, sdl2.RENDERER_ACCELERATED)
	if nil != err {
		return err
	}
	defer win.Destroy()

	screens := getScreens()
	scrID := SCR_MAIN_MENU
	ctrl := games.LoopCtrl{Quit: false, FPS: games.FPSControls{Capped: true, Number: FPS}}
	return games.MainLoop{
		Control: games.LoopCtrl{
			FPS: games.FPSControls{
				Capped: true,
				Number: 60,
			},
		},
	}.Launch(ctrl, func(loop games.LoopCtrl) (games.LoopCtrl, error) {
		scrID, err = screens[scrID].Execute(win)
		if nil != err {
			return games.LoopCtrl{Quit: true}, err
		}
		return games.LoopCtrl{Quit: SCR_NONE == scrID, FPS: loop.FPS}, nil
	})
}
