package aoc

import "testing"

func assertEqual(t *testing.T, expected int, got int) {
	if expected != got {
		t.Errorf("expected %d got %d", expected, got)
	}
}

func TestCalibrationValue(t *testing.T) {
	t.Run("1abc2", func(t *testing.T) {
		expected := 12
		got := CalibrationValue("1abc2")
		assertEqual(t, expected, got)
	})
}
