package cursor

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Mode uint8

const (
	CursorStatic Mode = iota
	CursorHide
)

type Model struct {
	Style     lipgloss.Style
	TextStyle lipgloss.Style
	char      string
	focus     bool
	mode      Mode
}

func New() Model {
	return Model{}
}

func (m *Model) Focus() {
	m.focus = true
	m.mode = CursorStatic
}

func (m *Model) Blur() {
	m.focus = false
	m.mode = CursorHide
}

func (m *Model) SetChar(char string) {
	m.char = char
}

func (m Model) Mode() Mode {
	return m.mode
}

func (m *Model) SetMode(mode Mode) {
	m.mode = mode
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	return m, nil
}

func (m Model) View() string {
	if m.mode == CursorHide || !m.focus {
		return ""
	}
	return m.Style.Inline(true).Reverse(true).Render(m.char)
}
