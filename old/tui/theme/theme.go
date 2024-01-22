package theme

type Variables map[string]string

type Rules map[string]string

type Theme struct {
	Variables       Variables       `yaml:"variables"`
	RootStyle       Style           `yaml:"root"`
	CommandBarTheme CommandBarTheme `yaml:"command_bar_theme"`
}

type CommandBarTheme struct {
	PrefixStyle     Style `yaml:"prefix_style"`
	TextStyle       Style `yaml:"text_style"`
	ErrCurrentStyle Style `yaml:"err_current_style"`
	ErrOldStyle     Style `yaml:"err_old_style"`
}

type Style struct {
	Rules Rules
}
