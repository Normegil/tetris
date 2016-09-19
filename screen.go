package main

import (
	"errors"
	"time"
)

type screen interface {
	execute(*window, *fonts) (ScreenID, error)
}

type DefaultScreen struct{}

func (d DefaultScreen) execute(*window, *fonts) (ScreenID, error) {
	return SCR_NONE, errors.New("Screen 'NONE' not meant to be used")
}

type ScreenID int

const (
	SCR_NONE        ScreenID = 0
	SCR_MAIN_MENU   ScreenID = 1
	SCR_EXIT_DIALOG ScreenID = 2
	SCR_PLAY                 = 3
)

func getScreens() map[ScreenID]screen {
	fpsCounter := &FPSCounter{lastTick: time.Now(), lastCalculated: time.Now()}

	screens := make(map[ScreenID]screen)
	screens[SCR_NONE] = DefaultScreen{}
	screens[SCR_EXIT_DIALOG] = &ExitScreen{
		counter: fpsCounter,
	}
	screens[SCR_MAIN_MENU] = &Menu{
		Title: "Tetris",
		items: []MenuItem{
			{Name: "New Game", ScrID: SCR_PLAY},
			{Name: "Options", ScrID: SCR_MAIN_MENU},
			{Name: "Quit", ScrID: SCR_EXIT_DIALOG},
		},
		counter: fpsCounter,
	}
	screens[SCR_PLAY] = &Play{
		counter: fpsCounter,
	}
	return screens
}
