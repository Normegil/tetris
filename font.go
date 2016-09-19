package main

import (
	"github.com/veandco/go-sdl2/sdl_ttf"
)

type fonts []font

func (f fonts) load(path string) *font {
	for _, font := range f {
		if path == font.path {
			return &font
		}
	}

	newFont := font{path: path}
	f = append(f, newFont)
	return &newFont
}

func (f fonts) Close() {
	for _, font := range f {
		font.Close()
	}
}

type font struct {
	path  string
	cache map[int]ttf.Font
}

func (f *font) size(size int) (ttf.Font, error) {
	if nil == f.cache {
		f.cache = make(map[int]ttf.Font)
	}

	_, present := f.cache[size]
	if !present {
		font, err := ttf.OpenFont(f.path, size)
		if nil != err {
			return ttf.Font{}, err
		}
		f.cache[size] = *font
	}

	return f.cache[size], nil
}

func (f *font) Close() {
	for _, fontToRelease := range f.cache {
		fontToRelease.Close()
	}
}
