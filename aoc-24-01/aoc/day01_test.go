package aoc

import "testing"

func assertEqual(t *testing.T, expected int, got int) {
	if expected != got {
		t.Errorf("expected %d got %d", expected, got)
	}
}

func assertEqualStr(t *testing.T, expected string, got string) {
	if expected != got {
		t.Errorf("expected %s got %s", expected, got)
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

	t.Run("treb7uchet", func(t *testing.T) {
		expected := 77
		got := CalibrationValue("treb7uchet")
		assertEqual(t, expected, got)
	})
}

func TestTotalCalibrationValue(t *testing.T) {
	t.Run("example1", func(t *testing.T) {
		input := `1abc2
		pqr3stu8vwx
		a1b2c3d4e5f
		treb7uchet`

		expected := 142
		got := TotalCalibrationValue(input)

		assertEqual(t, expected, got)
	})
}

func TestProcessLineForSpelledNumbers(t *testing.T) {
	t.Run("two1nine", func(t *testing.T) {
		expected := "219"
		got := ProcessLineForSpelledNumbers("two1nine")

		assertEqualStr(t, expected, got)
	})
}
