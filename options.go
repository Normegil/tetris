package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl"
)

type OptionsMenu struct {
	title    MenuMainTitle
	subTitle MenuSubTitle
	items    []MenuItem

	counter *FPSCounter
}

func NewOptionsMenu(counter *FPSCounter) *OptionsMenu {
	return &OptionsMenu{
		title:    MenuMainTitle{"Tetris"},
		subTitle: MenuSubTitle{"Options"},

		items: []MenuItem{
			&ScreenChangeMenuItem{BaseMenuItem: BaseMenuItem{Name: "Accept"}, ScreenID: SCR_MAIN_MENU},
			&ScreenChangeMenuItem{BaseMenuItem: BaseMenuItem{Name: "Cancel"}, ScreenID: SCR_MAIN_MENU},
		},

		counter: counter,
	}
}

func (s *OptionsMenu) Execute(win *window) (ScreenID, error) {
	scrID, err := s.handle(sdl.PollEvent())
	if nil != err {
		return SCR_NONE, err
	} else if SCR_OPTIONS != scrID {
		return scrID, nil
	}

	if err = win.Renderer().Clear(); nil != err {
		return SCR_NONE, err
	}

	if nil != s.counter {
		if err = s.counter.display(win.Renderer()); nil != err {
			return SCR_NONE, err
		}
	}

	if err = s.title.Render(win); nil != err {
		return SCR_NONE, err
	}

	subtitleY := int32(200)
	if err = s.subTitle.Render(win, s.title.Y()+subtitleY); nil != err {
		return SCR_NONE, err
	}

	itemsX := (win.GetSize().W - 400) / 2
	for i, item := range s.items {
		item.Render(win, sdl.Point{
			X: itemsX,
			Y: s.title.Y() + subtitleY + int32(100+i*80),
		})
	}

	win.Renderer().Present()

	return SCR_OPTIONS, nil
}

func (s *OptionsMenu) handle(ev sdl.Event) (ScreenID, error) {
	if nil != ev {
		switch ev.(type) {
		case *sdl.QuitEvent:
			logrus.Info("Quit event detected")
			return SCR_NONE, nil
		case *sdl.KeyDownEvent:
			keyDownEvent := ev.(*sdl.KeyDownEvent)
			switch keyDownEvent.Keysym.Sym {
			case sdl.K_ESCAPE:
				return SCR_MAIN_MENU, nil
			case sdl.K_UP:

			case sdl.K_DOWN:

			case sdl.K_KP_ENTER, sdl.K_RETURN:

			}
		}
	}
	return SCR_OPTIONS, nil
}
