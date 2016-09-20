package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type FPSCounter struct {
	lastTick       time.Time
	lastCalculated time.Time
	result         string
}

func NewCounter() *FPSCounter {
	return &FPSCounter{
		lastTick:       time.Now(),
		lastCalculated: time.Now(),
	}
}

func (s *FPSCounter) display(render *Renderer) error {
	now := time.Now()
	if int64(time.Second/4) < now.UnixNano()-s.lastCalculated.UnixNano() {
		s.lastCalculated = now
		timePassed := now.UnixNano() - s.lastTick.UnixNano()
		fps := float32(time.Second) / float32(timePassed)
		s.result = fmt.Sprintf("%.2g", fps)
	} else if "" == s.result {
		s.result = " "
	}
	err := render.Text(s.result, TextStyleWithPos{
		TextStyle: TextStyle{
			FontName: "capture-it.ttf",
			FontSize: 40,
			Color:    sdl.Color{R: 0, G: 255, B: 255, A: 255},
		},
		Position: sdl.Point{
			X: 10,
			Y: 10,
		},
	})
	s.lastTick = now
	return err
}
