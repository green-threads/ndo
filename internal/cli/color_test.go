package cli

import (
	"bytes"
	"os"
	"testing"

	"github.com/green-threads/ndo/internal/config"
)

func TestNewStylerDisabledForNonFileWriter(t *testing.T) {
	var buf bytes.Buffer
	s := newStyler(&buf, config.Settings{})
	if got := s.Bold("x"); got != "x" {
		t.Errorf("newStyler on a non-*os.File writer = colored, want passthrough (got %q)", got)
	}
}

func TestNewStylerDisabledForNonTerminalFile(t *testing.T) {
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()
	defer w.Close()

	// A pipe is a real *os.File but never an interactive terminal, so
	// color must stay off even though settings allow it.
	s := newStyler(w, config.Settings{})
	if got := s.Bold("x"); got != "x" {
		t.Errorf("newStyler on a pipe = colored, want passthrough (got %q)", got)
	}
}

func TestNewStylerRespectsExplicitColorFalse(t *testing.T) {
	falseVal := false
	var buf bytes.Buffer
	s := newStyler(&buf, config.Settings{Color: &falseVal})
	if got := s.Bold("x"); got != "x" {
		t.Errorf("newStyler with color=false = colored, want passthrough (got %q)", got)
	}
}

func TestColorAllowedNoColorEnvWinsOverSettings(t *testing.T) {
	t.Setenv("NO_COLOR", "1")
	trueVal := true
	if colorAllowed(config.Settings{Color: &trueVal}) {
		t.Error("colorAllowed() = true with NO_COLOR set and color=true, want false — NO_COLOR must win")
	}
}

func TestColorAllowedDefaultsTrue(t *testing.T) {
	if !colorAllowed(config.Settings{}) {
		t.Error("colorAllowed() = false for zero-value settings with no NO_COLOR, want true")
	}
}
