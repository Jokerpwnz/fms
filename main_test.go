package main

import (
	"testing"
)

func TestModThree(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		wantErr  bool
	}{
		{"1101", 1, false},
		{"1110", 2, false},
		{"1111", 0, false},
		{"110", 0, false},
		{"1010", 1, false},
		{"0", 0, false},
		{"1", 1, false},
		{"10", 2, false},
		{"11", 0, false},
		{"100", 1, false},
		{"", 0, true},          // Empty input
		{"102", 0, true},       // Invalid character
		{"1a0", 0, true},       // Invalid character
		{"1340", 0, true},       // Invalid character
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := modThree(tt.input)
			
			if tt.wantErr {
				if err == nil {
					t.Errorf("expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
			
			if result != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, result)
			}
		})
	}
}