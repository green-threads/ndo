package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/green-threads/ndo/internal/color"
	"github.com/green-threads/ndo/internal/config"
)

// PrintError writes "ndo: <err>" to deps.Stderr, in red when that stream
// is an interactive terminal and color isn't disabled. Called by main.go
// after root.Execute() returns a real ndo-level error (cobra's own error
// printing stays off via SilenceErrors, to avoid printing every error
// twice — see NewRootCmd). Settings are loaded best-effort here: if
// config.toml can't even be read, color still defaults on rather than
// hiding the original error behind a settings-loading one.
func PrintError(deps Deps, err error) {
	settings, loadErr := config.LoadSettings(deps.NDOHome)
	if loadErr != nil {
		settings = config.Settings{}
	}
	styler := newStyler(deps.Stderr, settings)
	fmt.Fprintln(deps.Stderr, styler.Red(fmt.Sprintf("ndo: %v", err)))
}

// newStyler decides whether w should receive ANSI-colored output.
// NO_COLOR (https://no-color.org) and an explicit `color = false` in
// config.toml both always disable it; otherwise color is used only when w
// is a real interactive terminal (never for pipes, files, or the
// in-memory writers tests use), enabling Windows' per-console VT100 mode
// first since that's off by default there.
func newStyler(w io.Writer, settings config.Settings) color.Styler {
	if !colorAllowed(settings) {
		return color.New(false)
	}
	f, ok := w.(*os.File)
	if !ok || !isTerminalFile(f) {
		return color.New(false)
	}
	return color.New(color.EnableVT(f))
}

// colorAllowed reports whether color is permitted before any terminal
// detection: NO_COLOR (https://no-color.org) always wins over settings,
// since it's a user's explicit, tool-wide opt-out.
func colorAllowed(settings config.Settings) bool {
	if _, noColor := os.LookupEnv("NO_COLOR"); noColor {
		return false
	}
	return settings.ColorEnabled()
}
