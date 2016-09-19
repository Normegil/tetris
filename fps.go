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

func (f *FPSCounter) display(render *renderer, fonts *fonts) error {
	const FONT = "capture-it.ttf"
	now := time.Now()
	if int64(time.Second/2) < now.UnixNano()-f.lastCalculated.UnixNano() {
		f.lastCalculated = now
		timePassed := now.UnixNano() - f.lastTick.UnixNano()
		fps := float32(time.Second) / float32(timePassed)
		f.result = fmt.Sprintf("%.2g", fps)
	} else if "" == f.result {
		f.result = " "
	}
	font, err := fonts.load(FONT).size(40)
	if nil != err {
		return err
	}
	text, err := render.Text(font, f.result, sdl.Color{R: 0, G: 255, B: 255, A: 255})
	if nil != err {
		return err
	}
	defer text.Close()

	err = render.Copy(text.Texture, &sdl.Rect{W: text.W, H: text.H}, &sdl.Rect{X: 5, Y: 5, W: text.W, H: text.H})
	if nil != err {
		return err
	}
	f.lastTick = now
	return nil
}
