package main

import (
	"testing"
	"time"
)

func TestParseDuration(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Duration
		wantErr  bool
	}{
		{"-2d", -48 * time.Hour, false},
		{"+3h", 3 * time.Hour, false},
		{"15m", 15 * time.Minute, false},
		{"10s", 10 * time.Second, false},
		{"", 0, false},
		{"bad", 0, true},
		{"5x", 0, true},
	}
	for _, tt := range tests {
		dur, err := parseDuration(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("parseDuration(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
		}
		if dur != tt.expected {
			t.Errorf("parseDuration(%q) = %v, want %v", tt.input, dur, tt.expected)
		}
	}
}

func TestFormatOutput(t *testing.T) {
    ts := time.Unix(1600000000, 0)
    if got := formatOutput(ts, true, false); got != "1600000000" {
        t.Errorf("formatOutput unix = %v, want 1600000000", got)
    }
    if got := formatOutput(ts, false, true); got != ts.UTC().Format(time.RFC3339) {
        t.Errorf("formatOutput utc = %v, want %v", got, ts.UTC().Format(time.RFC3339))
    }
    if got := formatOutput(ts, false, false); got != ts.Format(time.RFC3339) {
        t.Errorf("formatOutput default = %v, want %v", got, ts.Format(time.RFC3339))
    }
}

func TestFormatDuration(t *testing.T) {
	d := 49*time.Hour + 3*time.Minute + 7*time.Second
	want := "2d 1h 3m 7s"
	if got := formatDuration(d); got != want {
		t.Errorf("formatDuration = %v, want %v", got, want)
	}
}
