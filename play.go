package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/normegil/tetris/model"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	SQUARE_WIDTH  = 40
	SQUARE_HEIGHT = 40
)

type Play struct {
	tetris *model.Tetris

	counter *FPSCounter
}

func NewPlayScreen(counter *FPSCounter) *Play {
	return &Play{
		counter: counter,
		tetris:  model.NewTetris(),
	}
}

func (p *Play) Execute(win *window) (ScreenID, error) {
	scrID, err := p.handle(sdl.PollEvent())
	if nil != err {
		return SCR_NONE, err
	} else if SCR_PLAY != scrID {
		return scrID, nil
	}

	p.tetris.Update()
	lost := p.tetris.HasLost()

	if err = win.Renderer().Clear(); nil != err {
		return SCR_NONE, err
	}

	if lost {
		if err = p.displayLost(win); err != nil {
			return SCR_NONE, err
		}
	} else {
		if err = p.displayGame(win); err != nil {
			return SCR_NONE, err
		}
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

func (p Play) displayGame(win *window) error {
	err := p.displayLevel(win)
	if nil != err {
		return err
	}

	err = p.displayNextTetromino(win)
	if nil != err {
		return err
	}
	return nil
}

func (p Play) displayNextTetromino(win *window) error {
	rectCoord := sdl.Point{
		X: 2000,
		Y: 1200,
	}

	startPoint := sdl.Point{
		X: rectCoord.X + 2*(SQUARE_WIDTH+1),
		Y: rectCoord.Y + 2*(SQUARE_HEIGHT+1),
	}

	tetromino := p.tetris.NextTetromino
	for _, coordinate := range tetromino.AbsoluteCoordinates() {
		color := model.GetColor(tetromino.Type)
		point := sdl.Point{
			X: startPoint.X + int32(coordinate.X*(SQUARE_WIDTH+1)),
			Y: startPoint.Y + int32(coordinate.Y*(SQUARE_HEIGHT+1)),
		}
		err := p.displayBlock(win, point, sdl.Color{R: color.R, G: color.G, B: color.B, A: color.A})
		Once(func() {
			logrus.WithField("Rect Coord", point).WithField("Start", startPoint).Debug("Display Tetromino")
		})()
		if nil != err {
			return err
		}
	}

	return win.Renderer().CustomDrawColor(sdl.Color{255, 255, 255, 255}, func() error {
		rect := sdl.Rect{
			X: rectCoord.X,
			Y: rectCoord.Y,
			W: 6*(SQUARE_WIDTH+1) + 1,
			H: 6*(SQUARE_HEIGHT+1) + 1,
		}
		return win.Renderer().DrawRect(&rect)
	})
}

func (p Play) displayBlock(win *window, point sdl.Point, color sdl.Color) error {
	return win.Renderer().CustomDrawColor(color, func() error {
		return win.Renderer().FillRect(&sdl.Rect{
			X: point.X,
			Y: point.Y,
			W: SQUARE_WIDTH,
			H: SQUARE_HEIGHT,
		})
	})
}

func (p Play) displayLevel(win *window) error {
	msg := fmt.Sprintf("Level: %d", p.tetris.Level())
	return win.Renderer().Text(msg, TextStyleWithPos{
		TextStyle: TextStyle{
			FontName: FONT_TUSJ,
			FontSize: 30,
			Color:    sdl.Color{255, 255, 255, 255},
		},
		Position: sdl.Point{
			X: 2000,
			Y: 400,
		},
	})
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

func (p Play) squareSize() Size {
	return Size{
		W: 20,
		H: 20,
	}
}

var once bool

func Once(exec func()) func() {
	return func() {
		if !once {
			exec()
		}
		once = true
	}
}
