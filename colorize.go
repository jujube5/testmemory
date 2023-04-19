package testmemory

import (
	"fmt"
)

type Color string

const (
	RedBrightFont     = "\u001b[31;1m"
	GreenBrightFont   = "\u001b[32;1m"
	YellowBrightFont  = "\u001b[33;1m"
	BlueBrightFont    = "\u001b[34;1m"
	MagentaBrightFont = "\u001b[35;1m"
	CyanBrightFont    = "\u001b[36;1m"
	WhiteBrightFont   = "\u001b[37;1m"
	ColorReset        = "\u001b[0m"
)

func Colorize(color Color, message string) {
	fmt.Println(string(color), message, string(ColorReset))
}
