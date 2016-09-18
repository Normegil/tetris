package main

import "errors"

type screen func() (screen, error)

const (
	NONE screen = func() (screen, error){return nil, errors.New("Screen 'NONE' not meant to be used")}
	EXIT screen = exitScreen
	MENU screen = menuScreen
	PLAY screen = playScreen
)

