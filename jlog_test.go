package jlog

import (
	"testing"

	"github.com/Masterminds/log-go"
)

func TestEnabled(t *testing.T) {
	log.Current = New()
	if !Enabled() {
		t.Fatal("systemd journal is not available")
	}
}

func TestDebug(t *testing.T) {
	log.Current = New()
	log.Debug("Hello world")
}

func TestInfof(t *testing.T) {
	log.Current = New()
	world := "world"
	log.Infof("Hello %s", world)
}
