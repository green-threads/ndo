// Package color applies ANSI styling to strings, or passes them through
// unchanged when disabled — callers decide once whether color is on
// (settings, NO_COLOR, TTY, platform console mode) and then style
// unconditionally, rather than branching at every call site.
package color

import "strings"

// Styler wraps strings in ANSI escape codes when enabled.
type Styler struct {
	enabled bool
}

// New returns a Styler that applies styles only when enabled is true.
func New(enabled bool) Styler {
	return Styler{enabled: enabled}
}

// Bold returns str unchanged if disabled, otherwise bolded.
func (s Styler) Bold(str string) string {
	return s.style(str, "1")
}

// Dim returns str unchanged if disabled, otherwise dimmed.
func (s Styler) Dim(str string) string {
	return s.style(str, "2")
}

// Red returns str unchanged if disabled, otherwise red.
func (s Styler) Red(str string) string {
	return s.style(str, "31")
}

// BoldCyan returns str unchanged if disabled, otherwise bold cyan.
func (s Styler) BoldCyan(str string) string {
	return s.style(str, "1", "36")
}

func (s Styler) style(str string, codes ...string) string {
	if !s.enabled || str == "" {
		return str
	}
	return "\x1b[" + strings.Join(codes, ";") + "m" + str + "\x1b[0m"
}
