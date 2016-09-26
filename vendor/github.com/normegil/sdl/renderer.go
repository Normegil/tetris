package sdl

import (
	"github.com/veandco/go-sdl2/sdl"
)

// Renderer decorate an sdl renderer and change some method for easier and more consistent usage
type Renderer struct {
	*sdl.Renderer
	fonts *Fonts
}

// TextStyle define the style a text should have (Font, Size & Color)
type TextStyle struct {
	FontName string
	FontSize int
	Color    sdl.Color
}

// TextStyleWithPos define a text style as well as his position
type TextStyleWithPos struct {
	TextStyle
	Position sdl.Point
}

// NewRenderer construct an instance of Renderer
func NewRenderer(renderer *sdl.Renderer) *Renderer {
	return &Renderer{
		Renderer: renderer,
		fonts:    &Fonts{},
	}
}

// Text print a given text using the source renderer.
func (r *Renderer) Text(text string, style TextStyleWithPos) error {
	font, err := r.fonts.Load(style.FontName).Size(style.FontSize)
	if nil != err {
		return err
	}

	surface, err := font.RenderUTF8_Solid(text, style.Color)
	if nil != err {
		return err
	}
	defer surface.Free()

	t, err := r.CreateTextureFromSurface(surface)
	if nil != err {
		return err
	}
	defer t.Destroy()

	return r.Copy(t, &sdl.Rect{
		W: surface.W,
		H: surface.H,
	}, &sdl.Rect{
		X: style.Position.X,
		Y: style.Position.Y,
		W: surface.W,
		H: surface.H,
	})
}

// TextSize calculate the texture size the given text would use if it was printed
func (r *Renderer) TextSize(text string, style TextStyle) (Size, error) {
	font, err := r.fonts.Load(style.FontName).Size(style.FontSize)
	if nil != err {
		return Size{}, err
	}

	surface, err := font.RenderUTF8_Solid(text, style.Color)
	if nil != err {
		return Size{}, err
	}
	defer surface.Free()

	return Size{
		W: surface.W,
		H: surface.H,
	}, nil
}

// GetDrawColor redefine sdl.Renderer.GetDrawColor, sending an actual sdl.Color instance.
func (r Renderer) GetDrawColor() (sdl.Color, error) {
	red, g, b, a, err := r.Renderer.GetDrawColor()
	if nil != err {
		return sdl.Color{}, err
	}
	return sdl.Color{R: red, G: g, B: b, A: a}, nil
}

// SetDrawColor redefine sdl.Renderer.SetDrawColor, using sdl.Color
func (r *Renderer) SetDrawColor(color sdl.Color) error {
	return r.Renderer.SetDrawColor(color.R, color.G, color.B, color.A)
}

// DrawLine will draw a line between the 2 specified points of the given color
func (r *Renderer) DrawLine(color sdl.Color, source sdl.Point, target sdl.Point) error {
	return r.customDrawColor(color, func() error {
		return r.Renderer.DrawLine(int(source.X), int(source.Y), int(target.X), int(target.Y))
	})
}

// DrawLines will draw a line between the specified points of the given color
func (r *Renderer) DrawLines(color sdl.Color, points []sdl.Point) error {
	return r.customDrawColor(color, func() error {
		return r.Renderer.DrawLines(points)
	})
}

// DrawRect will draw a empty rectangle from the given go-sdl.Rect
func (r *Renderer) DrawRect(color sdl.Color, rect sdl.Rect) error {
	return r.customDrawColor(color, func() error {
		return r.Renderer.DrawRect(&rect)
	})
}

// FillRect will draw a filled rectangle from the given go-sdl.Rect
func (r *Renderer) FillRect(color sdl.Color, rect sdl.Rect) error {
	return r.customDrawColor(color, func() error {
		return r.Renderer.FillRect(&rect)
	})
}

func (r *Renderer) customDrawColor(color sdl.Color, toExec func() error) error {
	oldColor, err := r.GetDrawColor()
	if nil != err {
		return err
	}

	if err = r.SetDrawColor(color); nil != err {
		return err
	}

	if err = toExec(); nil != err {
		return err
	}

	return r.SetDrawColor(oldColor)
}

// Close the current renderer. If it's associated with a Window, use Window.Close()
func (r *Renderer) Close() {
	r.Renderer.Destroy()
}
