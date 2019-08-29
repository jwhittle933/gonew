package main

import (
	"log"
	"os"

	au "github.com/logrusorgru/aurora"
)

// Log ...
type Log struct {
	Trace,
	Info,
	Warning,
	Error *log.Logger
}

// StartLog ...
func StartLog() *Log {
	return &Log{
		Trace:   log.New(os.Stdout, au.Sprintf(au.BrightMagenta("[%s] "), "Trace"), log.Ldate|log.Ltime),
		Info:    log.New(os.Stdout, au.Sprintf(au.BrightGreen("[%s] "), "Info"), log.Ldate|log.Ltime),
		Warning: log.New(os.Stdout, au.Sprintf(au.BrightYellow("[%s] "), "Warning"), log.Ldate|log.Ltime),
		Error:   log.New(os.Stdout, au.Sprintf(au.BrightRed("[%s] "), "Error"), log.Ldate|log.Ltime),
	}
}

// T Trace method shorthand
func (l *Log) T(f string, m ...interface{}) {
	l.Trace.Printf(f, m...)
}

// I Info method shorthand
func (l *Log) I(f string, m ...interface{}) {
	l.Info.Printf(f, m...)
}

// W Warning method shorthand
func (l *Log) W(f string, m ...interface{}) {
	l.Warning.Printf(f, m...)
}

// E Error method shorthand
func (l *Log) E(f string, m ...interface{}) {
	l.Error.Printf(f, m...)
}
