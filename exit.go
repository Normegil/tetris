package main

import (
	"math/rand"

	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/normegil/sdl"
	"github.com/normegil/sdl/games"
	sdl2 "github.com/veandco/go-sdl2/sdl"
)

type ExitScreen struct {
	msg string

	counter games.FPSCounter
	fonts   *sdl.Fonts
}

func NewExitScreen(counter *games.UnlimitedFPSCounter) *ExitScreen {
	return &ExitScreen{
		counter: counter,
		fonts:   &sdl.Fonts{},
	}
}

func (e *ExitScreen) Execute(win *sdl.Window) (ScreenID, error) {
	scrID, err := e.handle(sdl2.PollEvent())
	if nil != err {
		return SCR_NONE, err
	} else if SCR_EXIT_DIALOG != scrID {
		logrus.WithField("Next Screen", scrID).Info("Changing screen")
		e.msg = ""
		return scrID, nil
	}

	err = win.Renderer().Clear()
	if nil != err {
		return SCR_NONE, err
	}

	if nil != e.counter {
		nbFps := e.counter.FPS()
		err = win.Renderer().Text(fmt.Sprintf("%g", nbFps), sdl.TextStyleWithPos{
			Position: sdl2.Point{
				X: 10,
				Y: 10,
			},
		})
		if nil != err {
			return SCR_NONE, err
		}
	}

	if err = e.displayMsg(win); nil != err {
		return SCR_NONE, err
	}

	win.Renderer().Present()
	return SCR_EXIT_DIALOG, nil
}

func (e *ExitScreen) displayMsg(win *sdl.Window) error {
	msgs := []string{
		"Too Hard for you ?",
		"Already Quitting ?",
		"Leaving us ?",
	}

	const FONT_SIZE = 100
	if "" == e.msg {
		e.msg = msgs[rand.Intn(len(msgs))]
	}

	exitTextStyle := sdl.TextStyle{
		FontName: FONT_TUSJ,
		FontSize: FONT_SIZE,
		Color:    sdl2.Color{R: 255, G: 255, B: 255, A: 255},
	}
	size, err := win.Renderer().TextSize(e.msg, exitTextStyle)
	if nil != err {
		return err
	}

	winSize := win.Size()
	msgY := (winSize.H - FONT_SIZE) / 2
	err = win.Renderer().Text(e.msg, sdl.TextStyleWithPos{
		TextStyle: exitTextStyle,
		Position: sdl2.Point{
			X: (winSize.W - size.W) / 2,
			Y: msgY,
		},
	})
	if nil != err {
		return err
	}

	choices := "( y / N )"
	exitTextStyle.FontSize = 80
	choiceSize, err := win.Renderer().TextSize(choices, exitTextStyle)
	if nil != err {
		return err
	}
	return win.Renderer().Text(choices, sdl.TextStyleWithPos{
		TextStyle: exitTextStyle,
		Position: sdl2.Point{
			X: (win.Size().W - choiceSize.W) / 2,
			Y: msgY + 120,
		},
	})
}

func (e *ExitScreen) handle(ev sdl2.Event) (ScreenID, error) {
	if nil != ev {
		switch ev.(type) {
		case *sdl2.QuitEvent:
			return SCR_NONE, nil
		case *sdl2.KeyDownEvent:
			keyDownEvent := ev.(*sdl2.KeyDownEvent)
			switch keyDownEvent.Keysym.Sym {
			case sdl2.K_y, sdl2.K_q:
				return SCR_NONE, nil
			case sdl2.K_n, sdl2.K_KP_ENTER, sdl2.K_RETURN:
				return SCR_MAIN_MENU, nil
			}
		}
	}
	return SCR_EXIT_DIALOG, nil
}
