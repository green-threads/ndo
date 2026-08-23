//go:build windows

package color

import (
	"os"

	"golang.org/x/sys/windows"
)

// EnableVT turns on ANSI escape processing for f's console. Windows
// Terminal already has this on, but legacy consoles (cmd.exe, older
// PowerShell hosts) don't — without it, raw escape codes print as visible
// garbage instead of colors. Returns false if f isn't a console this can
// be enabled on (e.g. redirected to a file or pipe), so callers fall back
// to plain output rather than emitting escape codes nothing will render.
func EnableVT(f *os.File) bool {
	handle := windows.Handle(f.Fd())
	var mode uint32
	if err := windows.GetConsoleMode(handle, &mode); err != nil {
		return false
	}
	mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	return windows.SetConsoleMode(handle, mode) == nil
}
