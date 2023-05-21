package adding

import "testing"

func TestStatusFormatStringToInt(t *testing.T) {
	tests := []struct {
		actual   string
		expected string
		want     int
	}{
		{"foo", "foo", 1},
		{"foo", "bar", 0},
	}
	for _, test := range tests {
		got := StatusFormatStringToInt(test.actual, test.expected)
		if got != test.want {
			t.Errorf("StatusFormatStringToInt(%v, %v) = %v, want %v", test.actual, test.expected, got, test.want)
		}
	}
}

func TestStatusFormatBoolToInt(t *testing.T) {
	tests := []struct {
		actual   bool
		expected bool
		want     int
	}{
		{true, true, 1},
		{false, false, 1},
		{false, true, 0},
		{true, false, 0},
	}
	for _, test := range tests {
		got := StatusFormatBoolToInt(test.actual, test.expected)
		if got != test.want {
			t.Errorf("StatusFormatBoolToInt(%v, %v) = %v, want %v", test.actual, test.expected, got, test.want)
		}
	}
}
