package sdl

import (
	"github.com/Sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

// Fonts load TTF fonts and put them in a cache
type Fonts struct {
	cache []*Font
}

// Load load the font found under given path
func (f *Fonts) Load(path string) *Font {
	for _, font := range f.cache {
		if path == font.path {
			return font
		}
	}

	newFont := &Font{path: path}
	f.cache = append(f.cache, newFont)
	return newFont
}

// Close font cache and all associated fonts
func (f *Fonts) Close() {
	for _, font := range f.cache {
		font.Close()
	}
}

// Font define a font with size cache
type Font struct {
	path  string
	cache map[int]*ttf.Font
}

// Size load the font with the specified size and cache the result
func (f *Font) Size(size int) (*ttf.Font, error) {
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

// Close the current font
func (f *Font) Close() {
	for _, fontToRelease := range f.cache {
		fontToRelease.Close()
	}
}
