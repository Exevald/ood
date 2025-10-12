package tests

import (
	"testing"

	"pictures/pkg/model"
)

func TestParseColor(t *testing.T) {
	tests := []struct {
		input   string
		want    model.Color
		wantErr bool
	}{
		{"red", model.Red, false},
		{"green", model.Green, false},
		{"blue", model.Blue, false},
		{"yellow", model.Yellow, false},
		{"pink", model.Pink, false},
		{"black", model.Black, false},
		{"invalid", nil, true},
		{"", nil, true},
	}

	for _, tt := range tests {
		got, err := model.ParseColor(tt.input)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseColor(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("ParseColor(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}
