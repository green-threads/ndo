package color

import "testing"

func TestStylerEnabled(t *testing.T) {
	s := New(true)
	tests := []struct {
		name string
		fn   func(string) string
		want string
	}{
		{"Bold", s.Bold, "\x1b[1mx\x1b[0m"},
		{"Dim", s.Dim, "\x1b[2mx\x1b[0m"},
		{"Red", s.Red, "\x1b[31mx\x1b[0m"},
		{"BoldCyan", s.BoldCyan, "\x1b[1;36mx\x1b[0m"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fn("x"); got != tt.want {
				t.Errorf("%s(%q) = %q, want %q", tt.name, "x", got, tt.want)
			}
		})
	}
}

func TestStylerDisabledIsPassthrough(t *testing.T) {
	s := New(false)
	for _, fn := range []func(string) string{s.Bold, s.Dim, s.Red, s.BoldCyan} {
		if got := fn("x"); got != "x" {
			t.Errorf("disabled styler = %q, want unchanged %q", got, "x")
		}
	}
}

func TestStylerEmptyStringUnchanged(t *testing.T) {
	s := New(true)
	if got := s.Bold(""); got != "" {
		t.Errorf("Bold(%q) = %q, want empty string with no escape codes", "", got)
	}
}
