package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

type fonts struct {
	cache []*font
}

func (f *fonts) load(path string) *font {
	for _, font := range f.cache {
		if path == font.path {
			return font
		}
	}

	newFont := &font{path: path}
	f.cache = append(f.cache, newFont)
	return newFont
}

func (f *fonts) Close() {
	for _, font := range f.cache {
		font.Close()
	}
}

type font struct {
	path  string
	cache map[int]*ttf.Font
}

func (f *font) size(size int) (*ttf.Font, error) {
	if nil == f.cache {
		f.cache = make(map[int]*ttf.Font)
	}

	_, present := f.cache[size]
	if !present {
		font, err := ttf.OpenFont(f.path, size)
		if nil != err {
			logrus.WithFields(logrus.Fields{
				"Path":  f.path,
				"Size":  size,
				"Cache": f.cache,
			}).Debug("Problem opening font")
			return nil, err
		}
		f.cache[size] = font
	}

	return f.cache[size], nil
}

func (f *font) Close() {
	for _, fontToRelease := range f.cache {
		fontToRelease.Close()
	}
}
