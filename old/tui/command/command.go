package command

import (
	"MicroPad/tui/common"
	"MicroPad/tui/cursor"
	"errors"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	focused    bool
	value      string
	err        error
	errCurrent bool
	Cursor     cursor.Model
	CursorPos  uint8
}

func New() Model {
	m := Model{
		focused: false,
		value:   "",
		Cursor:  cursor.New(),
	}

	m.Cursor.SetChar(" ")

	return m
}

func (m *Model) Focus() {
	m.focused = true
	m.Cursor.Focus()
}

func (m *Model) Blur() {
	m.focused = false
	m.Cursor.Blur()
}

func (m Model) Value() string {
	return m.value
}

func (m *Model) Clear() {
	m.err = nil
	m.errCurrent = false

	m.value = ""
}

func (m *Model) Execute() (tea.Cmd, error) {
	val := strings.Trim(m.value, " \n")

	m.errCurrent = false
	m.err = nil

	switch val {
	case "q", "quit":
		return tea.Quit, nil
	default:
		err := errors.New("unknown command")
		m.err = err
		m.errCurrent = true
		return nil, err
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m *Model) UpdateCursor() {
	if m.CursorPos == uint8(len(m.value))-1 {
		m.Cursor.SetChar(" ")
	}
	if len(m.value) != 0 && m.CursorPos < uint8(len(m.value))-1 {
		m.Cursor.SetChar(string(m.value[m.CursorPos+1]))
	}
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	if !m.focused {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()

		if key == "enter" {
			cmd, _ := m.Execute()
			return m, cmd
		}

		m.errCurrent = false

		if key == "backspace" {
			if len(m.value) > 0 {
				mByte := []byte(m.value)
				mByte = append(mByte[:max(0, m.CursorPos)], mByte[m.CursorPos+1:]...)
				m.value = string(mByte)
			}

			if len(m.value) > 0 && m.CursorPos != 0 {
				m.CursorPos--
			} else {
				m.UpdateCursor()
			}

			return m, nil
		}

		if key == ":" && len(m.value) == 0 {
			return m, nil
		}

		if key == "left" {
			if m.CursorPos > 0 {
				m.CursorPos--
			}
		}

		if key == "right" {
			if m.CursorPos < uint8(len(m.value))-1 && m.CursorPos < uint8(len(m.value)) {
				m.CursorPos++
			}
		}

		if key == "left" || key == "right" {
			m.UpdateCursor()
		}

		if key == "end" {
			m.CursorPos = uint8(len(m.value)) - 1
		}

		if strings.Contains(common.NormalChars, key) {
			if len(m.value) > 0 {
				m.CursorPos++
			}

			head := m.value[:m.CursorPos]
			tail := m.value[m.CursorPos:]
			m.value = head + key + tail
		}

	}

	return m, nil
}

func split(str string, index int) (string, string) {
	if index < 0 || index >= len(str) {
		return "", ""
	}
	return str[:index], str[index:]
}

func (m Model) View() string {
	val := m.value

	TextStyle := lipgloss.NewStyle().Copy()
	common.Theme.CommandBarTheme.TextStyle.Rules.ParseRules(TextStyle)

	PrefixStyle := lipgloss.NewStyle().Copy()
	common.Theme.CommandBarTheme.PrefixStyle.Rules.ParseRules(PrefixStyle)

	ErrCurrentStyle := lipgloss.NewStyle().Copy()
	common.Theme.CommandBarTheme.ErrCurrentStyle.Rules.ParseRules(ErrCurrentStyle)

	ErrOldStyle := lipgloss.NewStyle().Copy()
	common.Theme.CommandBarTheme.ErrOldStyle.Rules.ParseRules(ErrOldStyle)

	if len(m.value) >= 3 && m.CursorPos < uint8(len(m.value))-1 {
		s1, s2 := split(val, int(m.CursorPos)+1)
		_, s3 := split(s2, 1)
		val = s1 + m.Cursor.View() + s3
	}

	if len(m.value) == 2 && m.CursorPos == 0 {
		val = val[:1] + m.Cursor.View()
	}

	if m.CursorPos == uint8(len(m.value))-1 {
		val = val[:len(m.value)] + m.Cursor.View()
	}

	if len(m.value) == 0 {
		val = m.Cursor.View()
	}

	s := PrefixStyle.Render(":") + TextStyle.Render(val)

	if m.err != nil && m.errCurrent {
		s += ErrCurrentStyle.Render(m.err.Error())
	}

	if m.err != nil && !m.errCurrent {
		s += ErrOldStyle.Render(m.err.Error())
	}

	return TextStyle.Render(s)
}
