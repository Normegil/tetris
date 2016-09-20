package main

import (
	"github.com/normegil/tetris/model"
	"github.com/veandco/go-sdl2/sdl"
)

type Play struct {
	tetris model.Tetris

	counter *FPSCounter
}

func (p *Play) Execute(win *window) (ScreenID, error) {
	scrID, err := p.handle(sdl.PollEvent())
	if nil != err {
		return SCR_NONE, err
	} else if SCR_PLAY != scrID {
		return scrID, nil
	}

	p.tetris.Update()
	//lost := p.tetris.HasLost()

	if err = win.Renderer().Clear(); nil != err {
		return SCR_NONE, err
	}

	if err = p.displayLost(win); err != nil {
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

func (p Play) displayLost(win *window) error {
	msg := "Game Over"
	style := TextStyle{
		FontName: FONT_TUSJ,
		FontSize: 100,
		Color:    sdl.Color{255, 255, 255, 255},
	}

	size, err := win.Renderer().TextureSize(msg, style)
	if nil != err {
		return err
	}

	gameOverY := (win.GetSize().H - size.H) / 2
	err = win.Renderer().Text(msg, TextStyleWithPos{
		TextStyle: style,
		Position: sdl.Point{
			X: (win.GetSize().W - size.W) / 2,
			Y: gameOverY,
		},
	})
	if nil != err {
		return err
	}

	msg = "Push ESC to go to main menu"
	style = TextStyle{
		FontName: FONT_TUSJ,
		FontSize: 50,
		Color:    sdl.Color{255, 255, 255, 255},
	}

	size, err = win.Renderer().TextureSize(msg, style)
	if nil != err {
		return err
	}

	return win.Renderer().Text(msg, TextStyleWithPos{
		TextStyle: style,
		Position: sdl.Point{
			X: (win.GetSize().W - size.W) / 2,
			Y: gameOverY + 150,
		},
	})
}
