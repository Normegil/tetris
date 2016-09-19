package main

import (
	"math/rand"

	"github.com/Sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl"
)

type ExitScreen struct {
	msg string

	counter *FPSCounter
}

func (e *ExitScreen) execute(window *window, fonts *fonts) (ScreenID, error) {
	scrID, err := e.handle(sdl.PollEvent())
	if nil != err {
		return SCR_NONE, err
	} else if SCR_EXIT_DIALOG != scrID {
		logrus.WithField("Next Screen", scrID).Info("Changing screen")
		e.msg = ""
		return scrID, nil
	}

	render, err := window.Renderer()
	if nil != err {
		return SCR_NONE, err
	}
	err = render.Clear()
	if nil != err {
		return SCR_NONE, err
	}

	if nil != e.counter {
		if err = e.counter.display(render, fonts); nil != err {
			return SCR_NONE, err
		}
	}

	if err = e.displayMsg(window, fonts); nil != err {
		return SCR_NONE, err
	}

	render.Present()
	return SCR_EXIT_DIALOG, nil
}

func (e *ExitScreen) displayMsg(window *window, fonts *fonts) error {
	msgs := []string{
		"Too Hard for you ?",
		"Already Quitting ?",
		"Leaving us ?",
	}

	const FONT = "anudrg.ttf"
	const FONT_SIZE = 100
	font, err := fonts.load(FONT).size(FONT_SIZE)
	if nil != err {
		return err
	}

	render, err := window.Renderer()
	if nil != err {
		return err
	}

	color := sdl.Color{R: 255, G: 255, B: 255, A: 255}
	if "" == e.msg {
		e.msg = msgs[rand.Intn(len(msgs))]
	}
	msg, err := render.Text(font, e.msg, color)
	if nil != err {
		return err
	}
	defer msg.Close()

	winSize := window.GetSize()
	msgY := (winSize.H - FONT_SIZE) / 2
	err = render.Copy(
		msg.Texture,
		&sdl.Rect{W: msg.W, H: msg.H},
		&sdl.Rect{
			X: (winSize.W - msg.W) / 2,
			Y: msgY,
			W: msg.W,
			H: msg.H})
	if nil != err {
		return err
	}

	font, err = fonts.load(FONT).size(80)
	if nil != err {
		return err
	}
	possibilities, err := render.Text(font, "( y / N )", color)
	if nil != err {
		return err
	}
	defer possibilities.Close()

	return render.Copy(
		possibilities.Texture,
		&sdl.Rect{W: possibilities.W, H: possibilities.H},
		&sdl.Rect{
			X: (winSize.W - possibilities.W) / 2,
			Y: msgY + 120,
			W: possibilities.W,
			H: possibilities.H,
		},
	)
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
