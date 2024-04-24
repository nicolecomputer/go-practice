package aoc

import "testing"

func TestCalibrationValue(t *testing.T) {
	t.Run("1abc2", func(t *testing.T) {
		expected := 12
		got := CalibrationValue("1abc2")

		if expected != got {
			t.Errorf("expected %d got %d", expected, got)
		}
	})
}
