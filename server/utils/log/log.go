package log

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/kmiit/vivi/cmd/flags"
	"github.com/kmiit/vivi/types"
)

var logLevel = &flags.LogLevel

var (
	Fatal   = types.LogLevel{String: "FATAL", Index: 0, Color: color.RGB(174, 138, 190)}
	Error   = types.LogLevel{String: "ERROR", Index: 1, Color: color.RGB(245, 0, 0)}
	Warn    = types.LogLevel{String: "WARN", Index: 2, Color: color.RGB(187, 181, 41)}
	Info    = types.LogLevel{String: "INFO", Index: 3, Color: color.RGB(255, 255, 255)}
	Debug   = types.LogLevel{String: "DEBUG", Index: 4, Color: color.RGB(186, 186, 186)}
	Verbose = types.LogLevel{String: "VERBOSE", Index: 5, Color: color.RGB(128, 128, 128)}
)

func D(tag string, log ...any) {
	if *logLevel >= Debug.Index {
		logit(Debug, tag, log)
	}
}

func E(tag string, log ...any) {
	if *logLevel >= Error.Index {
		logit(Error, tag, log)
	}
}

func F(tag string, log ...any) {
	if *logLevel >= Fatal.Index {
		logit(Fatal, tag, log)
		os.Exit(1)
	}
}

func I(tag string, log ...any) {
	if *logLevel >= Info.Index {
		logit(Info, tag, log)
	}
}

func V(tag string, log ...any) {
	if *logLevel >= Verbose.Index {
		logit(Verbose, tag, log)
	}
}

func W(tag string, log ...any) {
	if *logLevel >= Warn.Index {
		logit(Warn, tag, log)
	}
}

func logit(loglevel types.LogLevel, tag string, log []any) {
	time := time.Now().Format("2006/01/02 15:04:05.00")
	prefix := fmt.Sprintf("[%s][%s]\t", loglevel.String, tag)
	content := fmt.Sprintln(log...)
	logline := fmt.Sprintf("%s %s %s", time, prefix, content)
	loglevel.Color.Printf("%s", logline)
}
