package main

import "errors"

type screen interface {
	execute(*window, *fonts) (screen, error)
	quit() bool
}

type DefaultScreen struct{}

func (d DefaultScreen) execute(*window, *fonts) (screen, error) {
	return nil, errors.New("Screen 'NONE' not meant to be used")
}

func (d DefaultScreen) quit() bool {
	return true
}
