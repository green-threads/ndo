//go:build !windows

package color

import "os"

// EnableVT is a no-op on non-Windows platforms — real terminals there
// already interpret ANSI escapes natively, with no console-mode opt-in.
func EnableVT(f *os.File) bool {
	return true
}
