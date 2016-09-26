package main

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/normegil/sdl"
	"github.com/normegil/sdl/games"
	sdl2 "github.com/veandco/go-sdl2/sdl"
)

type OptionsMenu struct {
	title    MenuMainTitle
	subTitle MenuSubTitle
	items    []MenuItem

	counter games.FPSCounter
}

func NewOptionsMenu(counter games.FPSCounter) *OptionsMenu {
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

func (s *OptionsMenu) Execute(win *sdl.Window) (ScreenID, error) {
	scrID, err := s.handle(sdl2.PollEvent())
	if nil != err {
		return SCR_NONE, err
	} else if SCR_OPTIONS != scrID {
		return scrID, nil
	}

	if err = win.Renderer().Clear(); nil != err {
		return SCR_NONE, err
	}

	if nil != s.counter {
		nbFps := s.counter.FPS()
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

	if err = s.title.Render(win); nil != err {
		return SCR_NONE, err
	}

	subtitleY := int32(200)
	if err = s.subTitle.Render(win, s.title.Y()+subtitleY); nil != err {
		return SCR_NONE, err
	}

	itemsX := (win.Size().W - 400) / 2
	for i, item := range s.items {
		item.Render(win, sdl2.Point{
			X: itemsX,
			Y: s.title.Y() + subtitleY + int32(100+i*80),
		})
	}

	win.Renderer().Present()

	return SCR_OPTIONS, nil
}

func (s *OptionsMenu) handle(ev sdl2.Event) (ScreenID, error) {
	if nil != ev {
		switch ev.(type) {
		case *sdl2.QuitEvent:
			logrus.Info("Quit event detected")
			return SCR_NONE, nil
		case *sdl2.KeyDownEvent:
			keyDownEvent := ev.(*sdl2.KeyDownEvent)
			switch keyDownEvent.Keysym.Sym {
			case sdl2.K_ESCAPE:
				return SCR_MAIN_MENU, nil
			case sdl2.K_UP:

			case sdl2.K_DOWN:

			case sdl2.K_KP_ENTER, sdl2.K_RETURN:

			}
		}
	}
	return SCR_OPTIONS, nil
}
