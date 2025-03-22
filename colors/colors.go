package colors

import (
	"github.com/fatih/color"
)

type ColorType = color.Color

var Magenta *color.Color = color.New(color.FgHiMagenta)
var Yellow *color.Color = color.New(color.FgHiYellow)
var Green *color.Color = color.New(color.FgGreen)
var Gray *color.Color = color.New(color.FgHiBlack)
var Blue *color.Color = color.New(color.FgBlue)
var Cyan *color.Color = color.New(color.FgCyan)
var Red *color.Color = color.New(color.FgHiRed)

var CyanBG *color.Color = color.New(color.BgCyan)
