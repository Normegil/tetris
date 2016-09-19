package main

import "github.com/veandco/go-sdl2/sdl"

type Play struct {
	counter *FPSCounter
}

func (p *Play) execute(window *window, fonts *fonts) (ScreenID, error) {
	scrID, err := p.handle(sdl.PollEvent())
	if nil != err {
		return SCR_NONE, err
	} else if SCR_PLAY != scrID {
		return scrID, nil
	}

	render, err := window.Renderer()
	if nil != err {
		return SCR_NONE, err
	}
	render.Clear()

	if nil != p.counter {
		if err = p.counter.display(render, fonts); nil != err {
			return SCR_NONE, err
		}
	}

	render.Present()
	return SCR_PLAY, nil
}

func (p *Play) handle(ev sdl.Event) (ScreenID, error) {
	if nil != ev {
		switch ev.(type) {
		case *sdl.QuitEvent:
			return SCR_NONE, nil
		case *sdl.KeyDownEvent:
			kdEvent := ev.(*sdl.KeyDownEvent)
			switch kdEvent.Keysym.Sym {
			case sdl.K_ESCAPE:
				return SCR_MAIN_MENU, nil
			}
		}
	}
	return SCR_PLAY, nil
}
