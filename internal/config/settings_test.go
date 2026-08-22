package config

import "testing"

func TestSettingsRoundTrip(t *testing.T) {
	dir := t.TempDir()
	falseVal := false
	want := Settings{Shell: "bash -c", Editor: "vim", Color: &falseVal, CompletionPromptAnswered: true}

	if err := SaveSettings(dir, want); err != nil {
		t.Fatalf("SaveSettings() error: %v", err)
	}
	got, err := LoadSettings(dir)
	if err != nil {
		t.Fatalf("LoadSettings() error: %v", err)
	}
	if got.Shell != want.Shell || got.Editor != want.Editor ||
		got.CompletionPromptAnswered != want.CompletionPromptAnswered ||
		got.ColorEnabled() != want.ColorEnabled() {
		t.Fatalf("LoadSettings() = %+v, want %+v", got, want)
	}
}

func TestColorEnabledDefaultsTrueWhenUnset(t *testing.T) {
	if !(Settings{}).ColorEnabled() {
		t.Fatal("ColorEnabled() = false for zero-value Settings, want true (unset defaults on)")
	}
}

func TestLoadSettingsMissingFileIsNotError(t *testing.T) {
	dir := t.TempDir()
	got, err := LoadSettings(dir)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != (Settings{}) {
		t.Fatalf("LoadSettings() = %+v, want zero value", got)
	}
}
