package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl"
)

type MenuItem struct {
	Name  string
	ScrID ScreenID
}

type Menu struct {
	Title    string
	items    []MenuItem
	selected int
}

func (m *Menu) execute(window *window, fonts *fonts) (ScreenID, error) {
	scrID, err := m.handle(sdl.PollEvent())
	if nil != err {
		return SCR_NONE, err
	} else if SCR_MAIN_MENU != scrID {
		logrus.WithField("Next Screen", scrID).Info("Changing screen")
		return scrID, nil
	}

	render, err := window.Renderer()
	if nil != err {
		return SCR_NONE, err
	}
	if err = render.Clear(); nil != err {
		return SCR_NONE, err
	}

	if err = m.renderGameTitle(window, fonts); nil != err {
		return SCR_NONE, err
	}

	if err = m.renderItems(window, fonts); nil != err {
		return SCR_NONE, err
	}

	render.Present()

	return SCR_MAIN_MENU, nil
}

func (m *Menu) handle(ev sdl.Event) (ScreenID, error) {
	if nil != ev {
		switch ev.(type) {
		case *sdl.QuitEvent:
			logrus.Info("Quit event detected")
			return SCR_NONE, nil
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
			case sdl.K_KP_ENTER, sdl.K_RIGHT, sdl.K_RETURN:
				return m.items[m.selected].ScrID, nil
			}
		}
	}
	return SCR_MAIN_MENU, nil
}

func (m *Menu) renderGameTitle(win *window, fonts *fonts) error {
	const (
		FT_GAME_TITLE      = "tusj.ttf"
		FT_SIZE_GAME_TITLE = 120
		TITLE_Y            = int32(200)
	)

	font, err := fonts.load(FT_GAME_TITLE).size(FT_SIZE_GAME_TITLE)
	if nil != err {
		return err
	}

	render, err := win.Renderer()
	if nil != err {
		return err
	}

	titleColor := sdl.Color{R: 255, G: 0, B: 0, A: 255}
	title, err := render.Text(font, m.Title, titleColor)
	if nil != err {
		return err
	}
	defer title.Close()

	winSize := win.GetSize()
	return render.Copy(
		title.Texture,
		&sdl.Rect{W: title.W, H: title.H},
		&sdl.Rect{
			X: (winSize.W - title.W) / 2,
			Y: TITLE_Y,
			W: title.W,
			H: title.H})
}

func (m *Menu) renderItems(win *window, fonts *fonts) error {
	const (
		FT_ITEM      = "anudrg.ttf"
		FT_SIZE_ITEM = 50
		ITEMS_Y      = int32(600)
		ITEMS_SPACE  = 100
	)

	itemFont, err := fonts.load(FT_ITEM).size(FT_SIZE_ITEM)
	if nil != err {
		return err
	}

	render, err := win.Renderer()
	if nil != err {
		return err
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
			return err
		}
		defer texture.Close()
		textures = append(textures, *texture)
	}

	itemsX := centerX(win.GetSize().W, textures)
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
			return err
		}
	}
	return nil
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
