package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	FT_GAME_TITLE      string = "tusj.ttf"
	FT_SIZE_GAME_TITLE        = 120
	TITLE_Y                   = int32(200)
	FT_ITEM            string = "anudrg.ttf"
	FT_SIZE_ITEM              = 50
	ITEMS_Y                   = int32(600)
	ITEMS_SPACE               = 100
)

type MenuItem struct {
	Name string
}

type Menu struct {
	Title    string
	items    []MenuItem
	selected int
}

func (m Menu) quit() bool {
	return false
}

func (m Menu) execute(window *window, fonts *fonts) (screen, error) {
	if ev := sdl.PollEvent(); nil != ev {
		switch ev.(type) {
		case *sdl.QuitEvent:
			logrus.Info("Quit event detected")
			return DefaultScreen{}, nil
		case *sdl.KeyDownEvent:
			keyDownEvent := ev.(*sdl.KeyDownEvent)
			switch keyDownEvent.Keysym.Sym {
			case sdl.K_UP:
				m.selected -= 1
				if m.selected < 0 {
					m.selected = len(m.items) - 1
				}
			case sdl.K_DOWN:
				m.selected += 1
				if m.selected >= len(m.items) {
					m.selected = 0
				}
			}
		}
	}

	gameTitleFont, err := fonts.load(FT_GAME_TITLE).size(FT_SIZE_GAME_TITLE)
	if nil != err {
		return DefaultScreen{}, err
	}

	itemFont, err := fonts.load(FT_ITEM).size(FT_SIZE_ITEM)
	if nil != err {
		return DefaultScreen{}, err
	}

	render, err := window.Renderer()
	if nil != err {
		return DefaultScreen{}, err
	}

	err = render.Clear()
	if nil != err {
		return DefaultScreen{}, err
	}

	titleColor := sdl.Color{R: 255, G: 0, B: 0, A: 255}
	title, err := render.Text(gameTitleFont, m.Title, titleColor)
	if nil != err {
		return DefaultScreen{}, err
	}
	defer title.Close()

	size := window.GetSize()
	err = render.Copy(
		title.Texture,
		&sdl.Rect{W: title.W, H: title.H},
		&sdl.Rect{
			X: (size.W - title.W) / 2,
			Y: TITLE_Y,
			W: title.W,
			H: title.H})
	if nil != err {
		return DefaultScreen{}, err
	}

	var textures []texture
	for i, item := range m.items {
		var color sdl.Color
		if m.selected == i {
			color = sdl.Color{R: 255, G: 255, B: 0, A: 255}
		} else {
			color = sdl.Color{R: 255, G: 255, B: 255, A: 255}
		}
		texture, err := render.Text(itemFont, item.Name, color)
		if nil != err {
			return DefaultScreen{}, err
		}
		defer texture.Close()
		textures = append(textures, *texture)
	}

	itemsX := centerX(size.W, textures)
	for i, texture := range textures {
		err = render.Copy(
			texture.Texture,
			&sdl.Rect{W: texture.W, H: texture.H},
			&sdl.Rect{
				X: itemsX,
				Y: ITEMS_Y + int32(i*ITEMS_SPACE),
				W: texture.W,
				H: texture.H})
		if nil != err {
			return DefaultScreen{}, err
		}
	}

	if nil != err {
		return DefaultScreen{}, err
	}
	render.Present()

	return m, nil
}

func centerX(WinWidth int32, textures []texture) int32 {
	var width int32
	for _, texture := range textures {
		if width < texture.W {
			width = texture.W
		}
	}

	return (WinWidth - width) / 2
}
