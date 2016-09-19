package main

import "github.com/veandco/go-sdl2/sdl"

const (
	FT_GAME_TITLE string = "FFF_Tusj.ttf"
	// FT_TITLE      string = ""
	FT_ITEM string = "anudi.ttf"
)

type MenuScreen struct{}

func (m MenuScreen) quit() bool {
	return false
}

func (m MenuScreen) execute(window *window, fonts *fonts) (screen, error) {
	if ev := sdl.PollEvent(); nil != ev {
		switch ev.(type) {
		case *sdl.QuitEvent:
			return DefaultScreen{}, nil
		}
	}

	gameTitleFont, err := fonts.load(FT_GAME_TITLE).size(30)
	if nil != err {
		return DefaultScreen{}, err
	}

	itemFont, err := fonts.load(FT_ITEM).size(12)
	if nil != err {
		return DefaultScreen{}, err
	}

	render, err := window.Renderer()
	if nil != err {
		return DefaultScreen{}, err
	}

	title, err := render.Text(gameTitleFont, "Tetris", sdl.Color{A: 255, R: 255, G: 0, B: 0})
	if nil != err {
		return DefaultScreen{}, err
	}
	defer title.Close()

	newGameItem, err := render.Text(itemFont, "New Game", sdl.Color{A: 255, R: 255, G: 255, B: 255})
	if nil != err {
		return DefaultScreen{}, err
	}
	defer title.Close()

	quitItem, err := render.Text(itemFont, "Quit", sdl.Color{A: 255, R: 255, G: 255, B: 255})
	if nil != err {
		return DefaultScreen{}, err
	}
	defer title.Close()

	err = render.Clear()
	if nil != err {
		return DefaultScreen{}, err
	}

	size := window.GetSize()
	err = render.Copy(title.Texture, &sdl.Rect{W: title.W, H: title.H}, &sdl.Rect{X: center(size.W, title.W), Y: 20, W: title.W, H: title.H})
	if nil != err {
		return DefaultScreen{}, err
	}
	err = render.Copy(newGameItem.Texture, &sdl.Rect{W: newGameItem.W, H: newGameItem.H}, &sdl.Rect{X: center(size.W, newGameItem.W), Y: 100, W: newGameItem.W, H: newGameItem.H})
	if nil != err {
		return DefaultScreen{}, err
	}
	err = render.Copy(quitItem.Texture, &sdl.Rect{W: quitItem.W, H: quitItem.H}, &sdl.Rect{X: center(size.W, quitItem.W), Y: 120, W: quitItem.W, H: quitItem.H})
	if nil != err {
		return DefaultScreen{}, err
	}
	render.Present()

	return m, nil
}

func center(windowW, textureW int32) int32 {
	return (windowW - textureW) / 2
}
