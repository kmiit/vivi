package log

import (
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/kmiit/vivi/types"
)

var (
	Fatal   = types.LogLevel{String: "FATAL", Index: 0, Color: color.RGB(174, 138, 190)}
	Error   = types.LogLevel{String: "ERROR", Index: 1, Color: color.RGB(245, 0, 0)}
	Warn    = types.LogLevel{String: "WARN", Index: 2, Color: color.RGB(187, 181, 41)}
	Info    = types.LogLevel{String: "INFO", Index: 3, Color: color.RGB(255, 255, 255)}
	Debug   = types.LogLevel{String: "DEBUG", Index: 4, Color: color.RGB(186, 186, 186)}
	Verbose = types.LogLevel{String: "VERBOSE", Index: 5, Color: color.RGB(128, 128, 128)}
)

func D(tag string, log ...any) {
	logit(Debug, tag, log)
}

func E(tag string, log ...any) {
	logit(Error, tag, log)
}

func F(tag string, log ...any) {
	logit(Fatal, tag, log)
	os.Exit(1)
}

func I(tag string, log ...any) {
	logit(Info, tag, log)
}

func V(tag string, log ...any) {
	logit(Verbose, tag, log)
}

func W(tag string, log ...any) {
	logit(Warn, tag, log)
}

func logit(loglevel types.LogLevel, tag string, log []any) {
	time := time.Now().Format("2006/01/02 15:04:05.00")
	prefix := fmt.Sprintf("[%s][%s]\t", loglevel.String, tag)
	content := fmt.Sprint(log...)
	logline := fmt.Sprintf("%s %s %s", time, prefix, content)
	loglevel.Color.Println(logline)
}
