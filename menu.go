package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl"
)

type MenuMainTitle struct {
	Title string
}

func (t MenuMainTitle) Render(win *window) error {
	const FONT_SIZE = 120

	titleStyle := TextStyle{
		FontName: FONT_ANUDRG,
		FontSize: FONT_SIZE,
		Color:    sdl.Color{R: 255, G: 0, B: 0, A: 255},
	}

	size, err := win.Renderer().TextureSize(t.Title, titleStyle)
	if nil != err {
		return err
	}

	return win.Renderer().Text(t.Title, TextStyleWithPos{
		TextStyle: titleStyle,
		Position: sdl.Point{
			X: (win.GetSize().W - size.W) / 2,
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

func (t MenuSubTitle) Render(win *window, y int32) error {
	const FONT_SIZE = 80

	style := TextStyle{
		FontName: FONT_TUSJ,
		FontSize: FONT_SIZE,
		Color:    sdl.Color{R: 125, G: 0, B: 125, A: 255},
	}
	size, err := win.Renderer().TextureSize(t.Title, style)
	if nil != err {
		return err
	}

	return win.Renderer().Text(t.Title, TextStyleWithPos{
		TextStyle: style,
		Position: sdl.Point{
			X: (win.GetSize().W - size.W) / 2,
			Y: y,
		},
	})
}

type MenuItem interface {
	Render(win *window, point sdl.Point) error
	Selected() bool
	Select()
	Unselect()
	Size(win *window) (Size, error)
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

func (i BaseMenuItem) style() TextStyle {
	var color sdl.Color
	if i.Selected() {
		color = sdl.Color{R: 255, G: 255, B: 0, A: 255}
	} else {
		color = sdl.Color{R: 255, G: 255, B: 255, A: 255}
	}

	return TextStyle{
		FontName: FONT_TUSJ,
		FontSize: 50,

		Color: color,
	}
}

func (i BaseMenuItem) Size(win *window) (Size, error) {
	return win.Renderer().TextureSize(i.Name, i.style())
}

func (i BaseMenuItem) Render(win *window, position sdl.Point) error {
	return win.Renderer().Text(i.Name, TextStyleWithPos{
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

	counter *FPSCounter
}

func NewMainMenu(counter *FPSCounter) *MainMenu {
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

func (m *MainMenu) Execute(win *window) (ScreenID, error) {
	scrID, err := m.handle(sdl.PollEvent())
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
		if err = m.counter.display(win.Renderer()); nil != err {
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
		point := sdl.Point{
			X: (win.GetSize().W - width) / 2,
			Y: m.title.Y() + int32(300+i*80),
		}
		if err := item.Render(win, point); nil != err {
			return SCR_NONE, err
		}
	}

	win.Renderer().Present()

	return SCR_MAIN_MENU, nil
}

func (m *MainMenu) handle(ev sdl.Event) (ScreenID, error) {
	if nil != ev {
		switch ev.(type) {
		case *sdl.QuitEvent:
			logrus.Info("Quit event detected")
			return SCR_NONE, nil
		case *sdl.KeyDownEvent:
			keyDownEvent := ev.(*sdl.KeyDownEvent)
			switch keyDownEvent.Keysym.Sym {
			case sdl.K_UP:
				m.SelectPrevious()
			case sdl.K_DOWN:
				m.SelectNext()
			case sdl.K_KP_ENTER, sdl.K_RIGHT, sdl.K_RETURN:
				_, item := m.Selected()
				return item.ScreenID, nil
			}
		}
	}
	return SCR_MAIN_MENU, nil
}

func (m *MainMenu) ItemsMaxWidth(win *window) (int32, error) {
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
