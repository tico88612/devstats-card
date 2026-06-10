package models

// Theme defines the color palette applied to a card.
type Theme struct {
	Background string
	Border     string
	TitleColor string
	TextColor  string
	IconColor  string
}

// DefaultTheme is the name used when no theme is requested or the name is unknown.
const DefaultTheme = "default"

// Themes maps a theme name (passed via the `theme` query parameter) to its palette.
var Themes = map[string]Theme{
	// Light themes
	"default": {
		Background: "#FFFEFE",
		Border:     "#E4E2E2",
		TitleColor: "#0086FF",
		TextColor:  "#555555",
		IconColor:  "#0086FF",
	},
	"solarized-light": {
		Background: "#FDF6E3",
		Border:     "#93A1A1",
		TitleColor: "#268BD2",
		TextColor:  "#586E75",
		IconColor:  "#B58900",
	},
	"gruvbox-light": {
		Background: "#FBF1C7",
		Border:     "#B57614",
		TitleColor: "#B57614",
		TextColor:  "#3C3836",
		IconColor:  "#D65D0E",
	},
	"latte": {
		Background: "#EFF1F5",
		Border:     "#CCD0DA",
		TitleColor: "#8839EF",
		TextColor:  "#4C4F69",
		IconColor:  "#1E66F5",
	},
	"vue": {
		Background: "#FFFEFE",
		Border:     "#41B883",
		TitleColor: "#41B883",
		TextColor:  "#273849",
		IconColor:  "#41B883",
	},
	// Dark themes
	"dark": {
		Background: "#0D1117",
		Border:     "#30363D",
		TitleColor: "#58A6FF",
		TextColor:  "#C9D1D9",
		IconColor:  "#58A6FF",
	},
	"dracula": {
		Background: "#282A36",
		Border:     "#6272A4",
		TitleColor: "#FF79C6",
		TextColor:  "#F8F8F2",
		IconColor:  "#BD93F9",
	},
	"radical": {
		Background: "#141321",
		Border:     "#FE428E",
		TitleColor: "#FE428E",
		TextColor:  "#A9FEF7",
		IconColor:  "#F8D847",
	},
	"gruvbox": {
		Background: "#282828",
		Border:     "#FABD2F",
		TitleColor: "#FABD2F",
		TextColor:  "#8EC07C",
		IconColor:  "#FE8019",
	},
	"tokyonight": {
		Background: "#1A1B27",
		Border:     "#70A5FD",
		TitleColor: "#70A5FD",
		TextColor:  "#38BDAE",
		IconColor:  "#BF91F3",
	},
}

// GetTheme returns the theme for name, falling back to the default theme when
// the name is unknown.
func GetTheme(name string) Theme {
	if theme, ok := Themes[name]; ok {
		return theme
	}
	return Themes[DefaultTheme]
}
