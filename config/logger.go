package config

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "", logger.Flags()),
		info:    log.New(writer, "", logger.Flags()),
		warning: log.New(writer, "", logger.Flags()),
		err:     log.New(writer, "", logger.Flags()),
		writer:  writer,
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println("DEBUG: " + fmt.Sprint(v...))
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println("INFO: " + fmt.Sprint(v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.warning.Println("WARNING: " + fmt.Sprint(v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println("ERROR: " + fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(colorCyan+"DEBUG: "+format+colorReset, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(colorGreen+"INFO: "+format+colorReset, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warning.Printf(colorYellow+"WARNING: "+format+colorReset, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(colorRed+"ERROR: "+format+colorReset, v...)
}
