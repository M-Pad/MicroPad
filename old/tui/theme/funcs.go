package theme

import "github.com/teacat/noire"

func Lighten(color string, factor float64) string {
	c := noire.NewHex(color)
	c = c.Lighten(factor)
	return "#" + c.Hex()
}

func Brighten(color string, factor float64) string {
	c := noire.NewHex(color)
	c = c.Brighten(factor)
	return "#" + c.Hex()
}

func Darken(color string, factor float64) string {
	c := noire.NewHex(color)
	c = c.Darken(factor)
	return "#" + c.Darken(factor).Hex()
}

func Shade(color string, factor float64) string {
	c := noire.NewHex(color)
	c = c.Shade(factor)
	return "#" + c.Hex()
}

func Saturate(color string, factor float64) string {
	c := noire.NewHex(color)
	c = c.Saturate(factor)
	return "#" + c.Hex()
}

func Desaturate(color string, factor float64) string {
	c := noire.NewHex(color)
	c = c.Desaturate(factor)
	return "#" + c.Hex()
}

func AdjustHue(color string, factor float64) string {
	c := noire.NewHex(color)
	c = c.AdjustHue(factor)
	return "#" + c.Hex()
}

func Mix(color1 string, color2 string, factor float64) string {
	c1 := noire.NewHex(color1)
	c2 := noire.NewHex(color2)
	return "#" + c1.Mix(c2, factor).Hex()
}

func Invert(color string) string {
	c := noire.NewHex(color)
	c = c.Invert()
	return "#" + c.Hex()
}

func Compliment(color string) string {
	c := noire.NewHex(color)
	c = c.Complement()
	return "#" + c.Hex()
}

func Grayscale(color string) string {
	c := noire.NewHex(color)
	c = c.Grayscale()
	return "#" + c.Hex()
}

func Foreground(color string) string {
	c := noire.NewHex(color)
	c = c.Foreground()
	return "#" + c.Hex()
}
