package main

import (
	"math/rand"

	"github.com/Sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl"
)

type ExitScreen struct {
	msg string

	counter *FPSCounter
	fonts   *fonts
}

func NewExitScreen(counter *FPSCounter) *ExitScreen {
	return &ExitScreen{
		counter: counter,
		fonts:   &fonts{},
	}
}

func (e *ExitScreen) Execute(win *window) (ScreenID, error) {
	scrID, err := e.handle(sdl.PollEvent())
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
		if err = e.counter.display(win.Renderer()); nil != err {
			return SCR_NONE, err
		}
	}

	if err = e.displayMsg(win); nil != err {
		return SCR_NONE, err
	}

	win.Renderer().Present()
	return SCR_EXIT_DIALOG, nil
}

func (e *ExitScreen) displayMsg(win *window) error {
	msgs := []string{
		"Too Hard for you ?",
		"Already Quitting ?",
		"Leaving us ?",
	}

	const FONT_SIZE = 100
	if "" == e.msg {
		e.msg = msgs[rand.Intn(len(msgs))]
	}

	exitTextStyle := TextStyle{
		FontName: FONT_TUSJ,
		FontSize: FONT_SIZE,
		Color:    sdl.Color{R: 255, G: 255, B: 255, A: 255},
	}
	size, err := win.Renderer().TextureSize(e.msg, exitTextStyle)
	if nil != err {
		return err
	}

	msgY := (win.GetSize().H - FONT_SIZE) / 2
	err = win.render.Text(e.msg, TextStyleWithPos{
		TextStyle: exitTextStyle,
		Position: sdl.Point{
			X: (win.GetSize().W - size.W) / 2,
			Y: msgY,
		},
	})
	if nil != err {
		return err
	}

	choices := "( y / N )"
	exitTextStyle.FontSize = 80
	choiceSize, err := win.Renderer().TextureSize(choices, exitTextStyle)
	if nil != err {
		return err
	}
	return win.Renderer().Text(choices, TextStyleWithPos{
		TextStyle: exitTextStyle,
		Position: sdl.Point{
			X: (win.GetSize().W - choiceSize.W) / 2,
			Y: msgY + 120,
		},
	})
}

func (e *ExitScreen) handle(ev sdl.Event) (ScreenID, error) {
	if nil != ev {
		switch ev.(type) {
		case *sdl.QuitEvent:
			return SCR_NONE, nil
		case *sdl.KeyDownEvent:
			keyDownEvent := ev.(*sdl.KeyDownEvent)
			switch keyDownEvent.Keysym.Sym {
			case sdl.K_y, sdl.K_q:
				return SCR_NONE, nil
			case sdl.K_n, sdl.K_KP_ENTER, sdl.K_RETURN:
				return SCR_MAIN_MENU, nil
			}
		}
	}
	return SCR_EXIT_DIALOG, nil
}
