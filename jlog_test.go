package jlog

import (
	"testing"

	"github.com/Masterminds/log-go"
)

func TestEnabled(t *testing.T) {
	log.Current = NewJournal(log.FatalLevel)
	if !Enabled() {
		t.Fatal("systemd journal is not available")
	}
}

func TestDebug(t *testing.T) {
	log.Current = NewJournal(log.DebugLevel)
	log.Debug("Hello world")
}

func TestInfof(t *testing.T) {
	log.Current = NewJournal(log.InfoLevel)
	world := "world"
	log.Infof("Hello %s", world)
}
