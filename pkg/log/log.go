package log

import (
	"io"
	"log"
	"os"
)

var (
	errorLog = log.New(os.Stdout, "\033[31m[error]\033[0m ", log.LstdFlags|log.Lshortfile)
	infoLog  = log.New(os.Stdout, "\033[34m[info]\033[0m ", log.LstdFlags|log.Lshortfile)
	debugLog = log.New(os.Stdout, "\033[32m[debug]\033[0m ", log.LstdFlags|log.Lshortfile)
	loggers  = []*log.Logger{errorLog, infoLog, debugLog}
)

var (
	Errorf = errorLog.Printf
	Infof  = infoLog.Printf
	Debugf = debugLog.Printf

	Error = errorLog.Println
	Info  = infoLog.Println
	Debug = debugLog.Println
)

const (
	Disabled = iota
	ErrorLevel
	InfoLevel
	DebugLevel
)

func SetLogLevel(level int) {
	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}
	switch level {
	case Disabled:
		errorLog.SetOutput(io.Discard)
		fallthrough
	case ErrorLevel:
		infoLog.SetOutput(io.Discard)
		fallthrough
	case InfoLevel:
		debugLog.SetOutput(io.Discard)
	}
	return
}
