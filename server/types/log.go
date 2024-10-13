package types

import (
	"github.com/fatih/color"
)

type LogLevel struct {
	String	string
	Index	int
	Color	*color.Color
}
