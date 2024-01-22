package tui

import (
	"MicroPad/tui/command"
	"MicroPad/tui/common"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Buffer struct {
	Name  string
	Path  string
	Saved bool
	Bytes []byte
}

type ViewMode uint8

const (
	ViewNormal  ViewMode = iota
	ViewInsert  ViewMode = iota
	ViewCommand ViewMode = iota
)

type Model struct {
	Buffers []Buffer
	Mode    ViewMode
	Key     string

	CommandBar command.Model
}

func (m Model) Init() tea.Cmd {
	return tea.EnterAltScreen
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "i":
			if m.Mode == ViewNormal {
				m.Mode = ViewInsert
			}
		case ":":
			if m.Mode == ViewNormal {
				m.Mode = ViewCommand
			}
		case "esc":
			if m.Mode != ViewNormal {
				m.Mode = ViewNormal
			}
		default:
			m.Key = msg.String()
		}
	}

	if m.Mode == ViewCommand {
		m.CommandBar.Focus()
	} else {
		m.CommandBar.Blur()
		m.CommandBar.Clear()
	}

	m.CommandBar, cmd = m.CommandBar.Update(msg)

	return m, cmd
}

func (m Model) View() string {
	s := ""

	RootStyle := lipgloss.NewStyle().Copy()
	common.Theme.RootStyle.Rules.ParseRules(RootStyle)

	switch m.Mode {
	case ViewMode(ViewNormal):
		s += "NORMAL"
	case ViewMode(ViewInsert):
		s += "INSERT"
	case ViewMode(ViewCommand):
		s += "COMMAND"
	}

	if m.Mode == ViewCommand {
		s += "\n" + m.CommandBar.View()
	}

	return RootStyle.Render(s)
}

func initialModel() Model {
	return Model{
		Buffers: []Buffer{
			{Name: "Untitled 1"},
		},
		CommandBar: command.New(),
	}
}

func TuiMain() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
