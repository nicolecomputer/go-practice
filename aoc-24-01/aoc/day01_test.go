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

	t.Run("pqr3stu8vwx", func(t *testing.T) {
		expected := 38
		got := CalibrationValue("pqr3stu8vwx")
		assertEqual(t, expected, got)
	})

	t.Run("a1b2c3d4e5f", func(t *testing.T) {
		expected := 15
		got := CalibrationValue("a1b2c3d4e5f")
		assertEqual(t, expected, got)
	})
}
