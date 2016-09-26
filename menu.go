package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/normegil/sdl"
	"github.com/normegil/sdl/games"
	sdl2 "github.com/veandco/go-sdl2/sdl"
)

type MenuMainTitle struct {
	Title string
}

func (t MenuMainTitle) Render(win *sdl.Window) error {
	const FONT_SIZE = 120

	titleStyle := sdl.TextStyle{
		FontName: FONT_ANUDRG,
		FontSize: FONT_SIZE,
		Color:    sdl2.Color{R: 255, G: 0, B: 0, A: 255},
	}

	size, err := win.Renderer().TextSize(t.Title, titleStyle)
	if nil != err {
		return err
	}

	return win.Renderer().Text(t.Title, sdl.TextStyleWithPos{
		TextStyle: titleStyle,
		Position: sdl2.Point{
			X: (win.Size().W - size.W) / 2,
			Y: t.Y(),
		},
	})
}

func (t MenuMainTitle) Y() int32 {
	return int32(200)
}

type MenuSubTitle struct {
	Title string
}

func (t MenuSubTitle) Render(win *sdl.Window, y int32) error {
	const FONT_SIZE = 80

	style := sdl.TextStyle{
		FontName: FONT_TUSJ,
		FontSize: FONT_SIZE,
		Color:    sdl2.Color{R: 125, G: 0, B: 125, A: 255},
	}
	size, err := win.Renderer().TextSize(t.Title, style)
	if nil != err {
		return err
	}

	return win.Renderer().Text(t.Title, sdl.TextStyleWithPos{
		TextStyle: style,
		Position: sdl2.Point{
			X: (win.Size().W - size.W) / 2,
			Y: y,
		},
	})
}

type MenuItem interface {
	Render(win *sdl.Window, point sdl2.Point) error
	Selected() bool
	Select()
	Unselect()
	Size(win *sdl.Window) (sdl.Size, error)
}

type BaseMenuItem struct {
	Name     string
	selected bool
}

func (i BaseMenuItem) Selected() bool {
	return i.selected
}

func (i *BaseMenuItem) Select() {
	i.selected = true
}

func (i *BaseMenuItem) Unselect() {
	i.selected = false
}

func (i BaseMenuItem) style() sdl.TextStyle {
	var color sdl2.Color
	if i.Selected() {
		color = sdl2.Color{R: 255, G: 255, B: 0, A: 255}
	} else {
		color = sdl2.Color{R: 255, G: 255, B: 255, A: 255}
	}

	return sdl.TextStyle{
		FontName: FONT_TUSJ,
		FontSize: 50,

		Color: color,
	}
}

func (i BaseMenuItem) Size(win *sdl.Window) (sdl.Size, error) {
	return win.Renderer().TextSize(i.Name, i.style())
}

func (i BaseMenuItem) Render(win *sdl.Window, position sdl2.Point) error {
	return win.Renderer().Text(i.Name, sdl.TextStyleWithPos{
		TextStyle: i.style(),
		Position:  position,
	})
}

type ScreenChangeMenuItem struct {
	BaseMenuItem
	ScreenID ScreenID
}

type MainMenu struct {
	title MenuMainTitle
	items []ScreenChangeMenuItem

	counter games.FPSCounter
}

func NewMainMenu(counter games.FPSCounter) *MainMenu {
	return &MainMenu{
		title: MenuMainTitle{"Tetris"},
		items: []ScreenChangeMenuItem{
			{BaseMenuItem: BaseMenuItem{Name: "New Game", selected: true}, ScreenID: SCR_PLAY},
			// {BaseMenuItem: BaseMenuItem{Name: "Options"}, ScreenID: SCR_OPTIONS},
			{BaseMenuItem: BaseMenuItem{Name: "Exit"}, ScreenID: SCR_EXIT_DIALOG},
		},
		counter: counter,
	}
}

func (m *MainMenu) Execute(win *sdl.Window) (ScreenID, error) {
	scrID, err := m.handle(sdl2.PollEvent())
	if nil != err {
		return SCR_NONE, err
	} else if SCR_MAIN_MENU != scrID {
		logrus.WithField("Next Screen", scrID).Info("Changing screen")
		return scrID, nil
	}

	if err = win.Renderer().Clear(); nil != err {
		return SCR_NONE, err
	}

	if nil != m.counter {
		nbFps := m.counter.FPS()
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

	if err = m.title.Render(win); nil != err {
		return SCR_NONE, err
	}

	width, err := m.ItemsMaxWidth(win)
	if nil != err {
		return SCR_NONE, err
	}

	for i, item := range m.items {
		point := sdl2.Point{
			X: (win.Size().W - width) / 2,
			Y: m.title.Y() + int32(300+i*80),
		}
		if err := item.Render(win, point); nil != err {
			return SCR_NONE, err
		}
	}

	win.Renderer().Present()

	return SCR_MAIN_MENU, nil
}

func (m *MainMenu) handle(ev sdl2.Event) (ScreenID, error) {
	if nil != ev {
		switch ev.(type) {
		case *sdl2.QuitEvent:
			logrus.Info("Quit event detected")
			return SCR_NONE, nil
		case *sdl2.KeyDownEvent:
			keyDownEvent := ev.(*sdl2.KeyDownEvent)
			switch keyDownEvent.Keysym.Sym {
			case sdl2.K_UP:
				m.SelectPrevious()
			case sdl2.K_DOWN:
				m.SelectNext()
			case sdl2.K_KP_ENTER, sdl2.K_RIGHT, sdl2.K_RETURN:
				_, item := m.Selected()
				return item.ScreenID, nil
			}
		}
	}
	return SCR_MAIN_MENU, nil
}

func (m *MainMenu) ItemsMaxWidth(win *sdl.Window) (int32, error) {
	var width int32

	for _, item := range m.items {
		size, err := item.Size(win)
		if nil != err {
			return 0, err
		}
		if width < size.W {
			width = size.W
		}
	}

	return width, nil
}

func (m *MainMenu) Selected() (int, *ScreenChangeMenuItem) {
	for i, item := range m.items {
		if item.Selected() {
			return i, &item
		}
	}
	return -1, &ScreenChangeMenuItem{ScreenID: SCR_MAIN_MENU}
}

func (m *MainMenu) SelectPrevious() {
	i, _ := m.Selected()
	if 0 > i {
		m.items[len(m.items)-1].Select()
	} else {
		m.items[i].Unselect()
		i -= 1
		if i < 0 {
			i = len(m.items) - 1
		}
		m.items[i].Select()
	}
}

func (m *MainMenu) SelectNext() {
	i, _ := m.Selected()
	if 0 > i {
		m.items[0].Select()
	} else {
		m.items[i].Unselect()
		i += 1
		if i >= len(m.items) {
			i = 0
		}
		m.items[i].Select()
	}
}
