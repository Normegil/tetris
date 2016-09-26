package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/normegil/sdl"
	"github.com/normegil/sdl/games"
	"github.com/normegil/tetris/model"
	sdl2 "github.com/veandco/go-sdl2/sdl"
)

const (
	SQUARE_WIDTH  = 60
	SQUARE_HEIGHT = 60
)

type Play struct {
	tetris *model.Tetris

	counter games.FPSCounter
}

func NewPlayScreen(counter games.FPSCounter) *Play {
	return &Play{
		counter: counter,
		tetris:  model.NewTetris(),
	}
}

func (p *Play) Execute(win *sdl.Window) (ScreenID, error) {
	scrID, err := p.handle(sdl2.PollEvent())
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
		nbFps := p.counter.FPS()
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

	win.Renderer().Present()
	return SCR_PLAY, nil
}

func (p *Play) handle(ev sdl2.Event) (ScreenID, error) {
	if nil != ev {
		switch ev.(type) {
		case *sdl2.QuitEvent:
			return SCR_NONE, nil
		case *sdl2.KeyDownEvent:
			kdEvent := ev.(*sdl2.KeyDownEvent)
			switch kdEvent.Keysym.Sym {
			case sdl2.K_ESCAPE:
				return SCR_MAIN_MENU, nil
			case sdl2.K_UP:
				p.tetris.Rotate(model.ROTATION_CLOCK)
			case sdl2.K_RIGHT:
				p.tetris.Move(model.DIRECTION_RIGHT)
			case sdl2.K_LEFT:
				p.tetris.Move(model.DIRECTION_LEFT)
			case sdl2.K_DOWN:
				p.tetris.Move(model.DIRECTION_DOWN)
			}
		}
	}
	return SCR_PLAY, nil
}

func (p Play) displayGame(win *sdl.Window) error {
	err := p.displayLevel(win)
	if nil != err {
		return err
	}

	err = p.displayNextTetromino(win)
	if nil != err {
		return err
	}
	return p.displayBoard(win)
}

func (p Play) displayBoard(win *sdl.Window) error {
	start := sdl2.Point{
		X: 1000,
		Y: 250,
	}

	Once(func() {
		logrus.Debug("Display Board")
	})()
	err := win.Renderer().DrawLines([]sdl2.Point{
		start,
		{X: start.X, Y: start.Y + 22*(SQUARE_HEIGHT+1)},
		{X: start.X + 10*(SQUARE_WIDTH+1) + 1, Y: start.Y + 22*(SQUARE_HEIGHT+1)},
		{X: start.X + 10*(SQUARE_WIDTH+1) + 1, Y: start.Y},
	}, sdl2.Color{255, 255, 255, 255})
	if nil != err {
		return err
	}

	err = win.Renderer().DrawLine(sdl2.Point{
		X: start.X,
		Y: start.Y + 2*(SQUARE_HEIGHT+1),
	}, sdl2.Point{
		X: start.X + 10*(SQUARE_WIDTH+1) + 1,
		Y: start.Y + 2*(SQUARE_HEIGHT+1),
	}, sdl2.Color{R: 255, A: 255})
	if nil != err {
		return err
	}

	tetromino := p.tetris.CurrentTetromino
	coord := tetromino.Center
	err = p.displayTetromino(win, tetromino, sdl2.Point{
		X: start.X + int32(1+coord.X*(SQUARE_WIDTH+1)),
		Y: start.Y + int32(1+coord.Y*(SQUARE_HEIGHT+1)),
	})
	if nil != err {
		return err
	}

	for _, square := range p.tetris.Board.Squares {
		err = p.displayBlock(win, sdl2.Point{
			X: start.X + int32(1+square.X*(SQUARE_WIDTH+1)),
			Y: start.Y + int32(1+square.Y*(SQUARE_HEIGHT+1)),
		}, sdl2.Color{
			R: square.Color.R,
			G: square.Color.G,
			B: square.Color.B,
			A: square.Color.A,
		})
		if nil != err {
			return err
		}
	}
	return nil
}

func (p Play) displayNextTetromino(win *sdl.Window) error {
	rectCoord := sdl2.Point{
		X: 2000,
		Y: 1200,
	}

	tetromino := p.tetris.NextTetromino
	err := p.displayTetromino(win, tetromino, sdl2.Point{
		X: rectCoord.X + 2*(SQUARE_WIDTH+1),
		Y: rectCoord.Y + 2*(SQUARE_HEIGHT+1),
	})
	if nil != err {
		return err
	}

	return win.Renderer().DrawRect(sdl2.Color{255, 255, 255, 255}, sdl2.Rect{
		X: rectCoord.X,
		Y: rectCoord.Y,
		W: 6*(SQUARE_WIDTH+1) + 1,
		H: 6*(SQUARE_HEIGHT+1) + 1,
	})
}

func (p Play) displayTetromino(win *sdl.Window, tetromino model.Tetromino, drawingCoordinates sdl2.Point) error {
	for _, coordinate := range tetromino.AbsoluteCoordinates() {
		color := model.GetColor(tetromino.Type)
		point := sdl2.Point{
			X: drawingCoordinates.X + int32(coordinate.X*(SQUARE_WIDTH+1)),
			Y: drawingCoordinates.Y + int32(coordinate.Y*(SQUARE_HEIGHT+1)),
		}
		err := p.displayBlock(win, point, sdl2.Color{R: color.R, G: color.G, B: color.B, A: color.A})
		if nil != err {
			return err
		}
	}
	return nil
}

func (p Play) displayBlock(win *sdl.Window, point sdl2.Point, color sdl2.Color) error {
	return win.Renderer().FillRect(&sdl2.Rect{
		X: point.X,
		Y: point.Y,
		W: SQUARE_WIDTH,
		H: SQUARE_HEIGHT,
	})
}

func (p Play) displayLevel(win *sdl.Window) error {
	msg := fmt.Sprintf("Level: %d", p.tetris.Level())
	return win.Renderer().Text(msg, sdl.TextStyleWithPos{
		TextStyle: sdl.TextStyle{
			FontName: FONT_TUSJ,
			FontSize: 30,
			Color:    sdl2.Color{255, 255, 255, 255},
		},
		Position: sdl2.Point{
			X: 2000,
			Y: 400,
		},
	})
}

func (p Play) displayLost(win *sdl.Window) error {
	msg := "Game Over"
	style := sdl.TextStyle{
		FontName: FONT_TUSJ,
		FontSize: 100,
		Color:    sdl2.Color{255, 255, 255, 255},
	}

	size, err := win.Renderer().TextSize(msg, style)
	if nil != err {
		return err
	}

	winSize := win.Size()
	gameOverY := (winSize.H - size.H) / 2
	err = win.Renderer().Text(msg, sdl.TextStyleWithPos{
		TextStyle: style,
		Position: sdl2.Point{
			X: (winSize.W - size.W) / 2,
			Y: gameOverY,
		},
	})
	if nil != err {
		return err
	}

	msg = "Push ESC to go to main menu"
	style = sdl.TextStyle{
		FontName: FONT_TUSJ,
		FontSize: 50,
		Color:    sdl2.Color{255, 255, 255, 255},
	}

	size, err = win.Renderer().TextSize(msg, style)
	if nil != err {
		return err
	}

	return win.Renderer().Text(msg, sdl.TextStyleWithPos{
		TextStyle: style,
		Position: sdl2.Point{
			X: (win.Size().W - size.W) / 2,
			Y: gameOverY + 150,
		},
	})
}

func (p Play) squareSize() sdl.Size {
	return sdl.Size{
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
