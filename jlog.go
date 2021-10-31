package jlog

import (
	"fmt"

	"github.com/Masterminds/log-go"
	"github.com/coreos/go-systemd/journal"
)

type Logger struct {
	logger jlogger
	Level  int
}

type jlogger struct{}

func New() *Logger {
	return &Logger{
		logger: jlogger{},
		Level:  log.InfoLevel,
	}
}

func Enabled() bool {
	return journal.Enabled()
}

func (l Logger) Trace(msg ...interface{}) {
	if l.Level <= log.TraceLevel {
		journal.Print(journal.PriDebug, fmt.Sprintf("%v", msg...))
	}
}

func (l Logger) Tracef(template string, args ...interface{}) {
	if l.Level <= log.TraceLevel {
		journal.Print(journal.PriDebug, template, args...)
	}
}

func (l Logger) Tracew(msg string, fields log.Fields) {
	if l.Level <= log.TraceLevel {
		journal.Send(msg, journal.PriDebug, handleFields(fields))
	}
}

func (l Logger) Debug(msg ...interface{}) {
	if l.Level <= log.DebugLevel {
		journal.Print(journal.PriDebug, fmt.Sprintf("%v", msg...))
	}
}

func (l Logger) Debugf(template string, args ...interface{}) {
	if l.Level <= log.TraceLevel {
		journal.Print(journal.PriDebug, template, args...)
	}
}

func (l Logger) Debugw(msg string, fields log.Fields) {
	if l.Level <= log.DebugLevel {
		journal.Send(msg, journal.PriDebug, handleFields(fields))
	}
}

func (l Logger) Info(msg ...interface{}) {
	if l.Level <= log.InfoLevel {
		journal.Print(journal.PriInfo, fmt.Sprintf("%v", msg...))
	}
}

func (l Logger) Infof(template string, args ...interface{}) {
	if l.Level <= log.InfoLevel {
		journal.Print(journal.PriInfo, template, args...)
	}
}

func (l Logger) Infow(msg string, fields log.Fields) {
	if l.Level <= log.InfoLevel {
		journal.Send(msg, journal.PriInfo, handleFields(fields))
	}
}
func (l Logger) Warn(msg ...interface{}) {
	if l.Level <= log.WarnLevel {
		journal.Print(journal.PriWarning, fmt.Sprintf("%v", msg...))
	}
}

func (l Logger) Warnf(template string, args ...interface{}) {
	if l.Level <= log.WarnLevel {
		journal.Print(journal.PriWarning, template, args...)
	}
}

func (l Logger) Warnw(msg string, fields log.Fields) {
	if l.Level <= log.WarnLevel {
		journal.Send(msg, journal.PriWarning, handleFields(fields))
	}
}

func (l Logger) Error(msg ...interface{}) {
	if l.Level <= log.ErrorLevel {
		journal.Print(journal.PriErr, fmt.Sprintf("%v", msg...))
	}
}

func (l Logger) Errorf(template string, args ...interface{}) {
	if l.Level <= log.ErrorLevel {
		journal.Print(journal.PriErr, template, args...)
	}
}

func (l Logger) Errorw(msg string, fields log.Fields) {
	if l.Level <= log.ErrorLevel {
		journal.Send(msg, journal.PriErr, handleFields(fields))
	}
}

func (l Logger) Fatal(msg ...interface{}) {
	if l.Level <= log.FatalLevel {
		journal.Print(journal.PriEmerg, fmt.Sprintf("%v", msg...))
	}
}

func (l Logger) Fatalf(template string, args ...interface{}) {
	if l.Level <= log.FatalLevel {
		journal.Print(journal.PriEmerg, template, args...)
	}
}

func (l Logger) Fatalw(msg string, fields log.Fields) {
	if l.Level <= log.FatalLevel {
		journal.Send(msg, journal.PriEmerg, handleFields(fields))
	}
}
func (l Logger) Panic(msg ...interface{}) {
	if l.Level <= log.PanicLevel {
		journal.Print(journal.PriEmerg, fmt.Sprintf("%v", msg...))
	}
}

func (l Logger) Panicf(template string, args ...interface{}) {
	if l.Level <= log.PanicLevel {
		journal.Print(journal.PriEmerg, template, args...)
	}
}

func (l Logger) Panicw(msg string, fields log.Fields) {
	if l.Level <= log.PanicLevel {
		journal.Send(msg, journal.PriEmerg, handleFields(fields))
	}
}

func handleFields(flds log.Fields) map[string]string {
	m := make(map[string]string, len(flds))
	for k, v := range flds {
		m[k] = fmt.Sprintf("%v", v)
	}
	return m
}
