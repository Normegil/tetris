package main

import "github.com/veandco/go-sdl2/sdl"

type Play struct {
	counter *FPSCounter
}

func (p *Play) Execute(win *window) (ScreenID, error) {
	scrID, err := p.handle(sdl.PollEvent())
	if nil != err {
		return SCR_NONE, err
	} else if SCR_PLAY != scrID {
		return scrID, nil
	}

	if err = win.Renderer().Clear(); nil != err {
		return SCR_NONE, err
	}

	if nil != p.counter {
		if err = p.counter.display(win.Renderer()); nil != err {
			return SCR_NONE, err
		}
	}

	win.Renderer().Present()
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
