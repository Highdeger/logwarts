package logwarts

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	instance *log.Logger
	verbose  bool
)

type Type string

const (
	TypeVerbose Type = "verbose"
	TypeInfo    Type = "info"
	TypeDebug   Type = "debug"
	TypeError   Type = "error"
	TypeFatal   Type = "fatal"
)

func (t Type) String() string {
	return strings.ToUpper(string(t))
}

func get() *log.Logger {
	if instance == nil {
		instance = log.Default()
		instance.SetOutput(os.Stdout)
		instance.SetPrefix("")
		instance.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	}
	return instance
}

func rawLog(logType Type, message, detail string, printStack bool) {
	if logType == TypeVerbose && !verbose {
		return
	}
	text := fmt.Sprintf("<%s> %s: %s", logType, message, detail)
	if printStack {
		text += fmt.Sprintf("\n%s", stack())
	}
	get().Print(text)
}

func Verbose(message, detail string) {
	rawLog(TypeVerbose, message, detail, false)
}

func VerboseStack(message, detail string) {
	rawLog(TypeVerbose, message, detail, true)
}

func Info(message, detail string) {
	rawLog(TypeInfo, message, detail, false)
}

func InfoStack(message, detail string) {
	rawLog(TypeInfo, message, detail, true)
}

func Debug(message, detail string) {
	rawLog(TypeDebug, message, detail, false)
}

func DebugStack(message, detail string) {
	rawLog(TypeDebug, message, detail, true)
}

func Error(message, detail string) {
	rawLog(TypeError, message, detail, false)
}

func ErrorStack(message, detail string) {
	rawLog(TypeError, message, detail, true)
}

func Fatal(message, detail string) {
	rawLog(TypeFatal, message, detail, false)
	os.Exit(1)
}

func FatalStack(message, detail string) {
	rawLog(TypeFatal, message, detail, true)
	os.Exit(1)
}
