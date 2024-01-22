package theme

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"text/template"

	"github.com/charmbracelet/lipgloss"
	"gopkg.in/yaml.v3"
)

func parsePosition(position string) (lipgloss.Position, error) {
	flt, err := strconv.ParseFloat(position, 64)

	return lipgloss.Position(flt), err
}

func parseColor(color string) lipgloss.TerminalColor {
	return lipgloss.Color(color)
}

func parseBool(val string) (bool, error) {
	return strconv.ParseBool(val)
}

func parseJSON(val string, obj any) error {
	return json.Unmarshal([]byte(val), obj)
}

func parseInt(val string) (int, error) {
	i64, err := strconv.ParseInt(val, 10, strconv.IntSize)
	return int(i64), err
}

func (rules Rules) ParseRules(style lipgloss.Style) error {
	for rule, value := range rules {
		switch rule {
		case "AlignHorizontal":
			pos, err := parsePosition(value)
			if err != nil {
				return err
			}
			style.AlignHorizontal(pos)
		case "AlignVertical":
			pos, err := parsePosition(value)
			if err != nil {
				return err
			}
			style.AlignVertical(pos)
		case "Background":
			style.Background(parseColor(value))
		case "Blink":
			bol, err := parseBool(value)
			if err != nil {
				return err
			}
			style.Blink(bol)
		case "Bold":
			bol, err := parseBool(value)
			if err != nil {
				return err
			}
			style.Bold(bol)
		case "BorderBackground":
			style.BorderBackground(parseColor(value))
		case "BorderBottom":
			bol, err := parseBool(value)
			if err != nil {
				return err
			}
			style.BorderBottom(bol)
		case "BorderBottomBackground":
			style.BorderBottomBackground(parseColor(value))
		case "BorderBottomForeground":
			style.BorderBottomForeground(parseColor(value))
		case "BorderForeground":
			style.BorderForeground(parseColor(value))
		case "BorderLeft":
			bol, err := parseBool(value)
			if err != nil {
				return err
			}
			style.BorderLeft(bol)
		case "BorderLeftBackground":
			style.BorderLeftBackground(parseColor(value))
		case "BorderLeftForeground":
			style.BorderLeftForeground(parseColor(value))
		case "BorderRight":
			bol, err := parseBool(value)
			if err != nil {
				return err
			}
			style.BorderRight(bol)
		case "BorderRightBackground":
			style.BorderRightBackground(parseColor(value))
		case "BorderRightForeground":
			style.BorderRightForeground(parseColor(value))
		case "BorderStyle":
			borderStyle := lipgloss.Border{}
			if err := parseJSON(value, &borderStyle); err != nil {
				return err
			}
		case "BorderTop":
			bol, err := parseBool(value)
			if err != nil {
				return err
			}
			style.BorderTop(bol)
		case "BorderTopBackground":
			style.BorderTopBackground(parseColor(value))
		case "BorderTopForeground":
			style.BorderTopForeground(parseColor(value))
		case "ColorWhitespace":
			bol, err := parseBool(value)
			if err != nil {
				return err
			}
			style.ColorWhitespace(bol)
		case "Faint":
			bol, err := parseBool(value)
			if err != nil {
				return err
			}
			style.Faint(bol)
		case "Foreground":
			style.Foreground(parseColor(value))
		case "Height":
			i, err := parseInt(value)
			if err != nil {
				return err
			}
			style.Height(i)
		case "Inline":
			bol, err := parseBool(value)
			if err != nil {
				return err
			}
			style.Inline(bol)
		case "Italic":
			bol, err := parseBool(value)
			if err != nil {
				return err
			}
			style.Italic(bol)
		case "MarginBackground":
			style.MarginBackground(parseColor(value))
		case "MarginBottom":
			i, err := parseInt(value)
			if err != nil {
				return err
			}
			style.MarginBottom(i)
		case "MarginLeft":
			i, err := parseInt(value)
			if err != nil {
				return err
			}
			style.MarginLeft(i)
		case "MarginRight":
			i, err := parseInt(value)
			if err != nil {
				return err
			}
			style.MarginRight(i)
		case "MarginTop":
			i, err := parseInt(value)
			if err != nil {
				return err
			}
			style.MarginTop(i)
		case "MaxHeight":
			i, err := parseInt(value)
			if err != nil {
				return err
			}
			style.MaxHeight(i)
		case "MaxWidth":
			i, err := parseInt(value)
			if err != nil {
				return err
			}
			style.MaxWidth(i)
		case "PaddingBottom":
			i, err := parseInt(value)
			if err != nil {
				return err
			}
			style.PaddingBottom(i)
		case "PaddingLeft":
			i, err := parseInt(value)
			if err != nil {
				return err
			}
			style.PaddingLeft(i)
		case "PaddingRight":
			i, err := parseInt(value)
			if err != nil {
				return err
			}
			style.PaddingRight(i)
		case "PaddingTop":
			i, err := parseInt(value)
			if err != nil {
				return err
			}
			style.PaddingTop(i)
		case "Reverse":
			bol, err := parseBool(value)
			if err != nil {
				return err
			}
			style.Reverse(bol)
		case "Strikethrough":
			bol, err := parseBool(value)
			if err != nil {
				return err
			}
			style.Strikethrough(bol)
		case "StrikethroughSpaces":
			bol, err := parseBool(value)
			if err != nil {
				return err
			}
			style.StrikethroughSpaces(bol)
		case "TabWidth":
			i, err := parseInt(value)
			if err != nil {
				return err
			}
			style.TabWidth(i)
		case "Underline":
			bol, err := parseBool(value)
			if err != nil {
				return err
			}
			style.Underline(bol)
		case "UnderlineSpaces":
			bol, err := parseBool(value)
			if err != nil {
				return err
			}
			style.UnderlineSpaces(bol)
		default:
			return fmt.Errorf("unknown rule: %s", rule)
		}
	}

	return nil
}

func LoadFromPath(path string) (Theme, error) {
	style := Theme{}

	byteArr, err := os.ReadFile(path)
	if err != nil {
		return style, err
	}

	if err := yaml.Unmarshal(byteArr, &style); err != nil {
		return style, err
	}

	buf := bytes.NewBuffer([]byte{})

	tmpl := template.New("tmpl")

	tmpl.Funcs(template.FuncMap{
		"lighten":    Lighten,
		"brighten":   Brighten,
		"darken":     Darken,
		"shade":      Shade,
		"saturate":   Saturate,
		"desaturate": Desaturate,
		"adjust_hue": AdjustHue,
		"invert":     Invert,
		"compliment": Compliment,
		"foreground": Foreground,
	})

	tmpl, err = tmpl.Parse(string(byteArr))
	if err != nil {
		return style, err
	}

	tmpl.Execute(buf, style.Variables)

	if err := yaml.Unmarshal(buf.Bytes(), &style); err != nil {
		return style, err
	}
	return style, nil
}

func SaveTheme(path string, theme Theme) error {
	b, err := yaml.Marshal(&theme)
	if err != nil {
		return err
	}

	return os.WriteFile("./test.yml", b, 0644)
}
