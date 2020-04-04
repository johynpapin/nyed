package config

import (
	"github.com/gdamore/tcell"
)

type ThemeConfig struct {
	Background string
	Foreground string
	Function   string
	Type       string
	Property   string
	Variable   string
	Operator   string
	Keyword    string
	String     string
	Escape     string
	Number     string
	Constant   string
	Comment    string
}

type Theme struct {
	config         *ThemeConfig
	highlightNames map[string]tcell.Color
	Style          tcell.Style
	Foreground     tcell.Color
	Background     tcell.Color
}

func NewTheme(themeConfig *ThemeConfig) *Theme {
	theme := &Theme{
		config:         themeConfig,
		highlightNames: make(map[string]tcell.Color),

		Foreground: tcell.GetColor(themeConfig.Foreground),
		Background: tcell.GetColor(themeConfig.Background),
	}

	theme.Style = tcell.StyleDefault.Background(theme.Background).Foreground(theme.Foreground)

	theme.highlightNames["function"] = tcell.GetColor(themeConfig.Function)
	theme.highlightNames["type"] = tcell.GetColor(themeConfig.Type)
	theme.highlightNames["property"] = tcell.GetColor(themeConfig.Property)
	theme.highlightNames["variable"] = tcell.GetColor(themeConfig.Variable)
	theme.highlightNames["operator"] = tcell.GetColor(themeConfig.Operator)
	theme.highlightNames["keyword"] = tcell.GetColor(themeConfig.Keyword)
	theme.highlightNames["string"] = tcell.GetColor(themeConfig.String)
	theme.highlightNames["escape"] = tcell.GetColor(themeConfig.Escape)
	theme.highlightNames["number"] = tcell.GetColor(themeConfig.Number)
	theme.highlightNames["constant"] = tcell.GetColor(themeConfig.Constant)
	theme.highlightNames["constant.builtin"] = tcell.GetColor(themeConfig.Constant)
	theme.highlightNames["comment"] = tcell.GetColor(themeConfig.Comment)

	return theme
}

func (theme *Theme) ColorFromCaptureName(captureName string) tcell.Color {
	color, exists := theme.highlightNames[captureName]
	if !exists || color == tcell.ColorDefault {
		return theme.Foreground
	}

	return color
}
