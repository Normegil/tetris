package main

import "errors"

type screen interface {
	Execute(*window) (ScreenID, error)
}

type DefaultScreen struct{}

func (d DefaultScreen) Execute(*window) (ScreenID, error) {
	return SCR_NONE, errors.New("Screen 'NONE' not meant to be used")
}

type ScreenID int

const (
	SCR_NONE        ScreenID = 0
	SCR_MAIN_MENU   ScreenID = 1
	SCR_EXIT_DIALOG ScreenID = 2
	SCR_PLAY        ScreenID = 3
	SCR_OPTIONS     ScreenID = 4
)

func getScreens() map[ScreenID]screen {
	fpsCounter := NewCounter()

	screens := make(map[ScreenID]screen)
	screens[SCR_NONE] = DefaultScreen{}
	screens[SCR_EXIT_DIALOG] = NewExitScreen(fpsCounter)
	screens[SCR_MAIN_MENU] = NewMainMenu(fpsCounter)
	screens[SCR_PLAY] = NewPlayScreen(fpsCounter)
	screens[SCR_OPTIONS] = NewOptionsMenu(fpsCounter)
	return screens
}
