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
	t.Run("1abc2 without filtering spelled words", func(t *testing.T) {
		expected := 12
		got := CalibrationValue("1abc2", false)
		assertEqual(t, expected, got)
	})

	t.Run("pqr3stu8vwx without filtering spelled words", func(t *testing.T) {
		expected := 38
		got := CalibrationValue("pqr3stu8vwx", false)
		assertEqual(t, expected, got)
	})

	t.Run("a1b2c3d4e5f without filtering spelled words", func(t *testing.T) {
		expected := 15
		got := CalibrationValue("a1b2c3d4e5f", false)
		assertEqual(t, expected, got)
	})

	t.Run("treb7uchet without filtering spelled words", func(t *testing.T) {
		expected := 77
		got := CalibrationValue("treb7uchet", false)
		assertEqual(t, expected, got)
	})

	t.Run("two1nine with filtering spelled words", func(t *testing.T) {
		expected := 29
		got := CalibrationValue("two1nine", true)
		assertEqual(t, expected, got)
	})

	t.Run("eightwothree with filtering spelled words", func(t *testing.T) {
		expected := 83
		got := CalibrationValue("eightwothree", true)
		assertEqual(t, expected, got)
	})

	t.Run("abcone2threexyz with filtering spelled words", func(t *testing.T) {
		expected := 13
		got := CalibrationValue("abcone2threexyz", true)
		assertEqual(t, expected, got)
	})

	t.Run("xtwone3four with filtering spelled words", func(t *testing.T) {
		expected := 24
		got := CalibrationValue("xtwone3four", true)
		assertEqual(t, expected, got)
	})

	t.Run("4nineeightseven2 with filtering spelled words", func(t *testing.T) {
		expected := 42
		got := CalibrationValue("4nineeightseven2", true)
		assertEqual(t, expected, got)
	})

	t.Run("zoneight234 with filtering spelled words", func(t *testing.T) {
		expected := 14
		got := CalibrationValue("zoneight234", true)
		assertEqual(t, expected, got)
	})

	t.Run("7pqrstsixteen with filtering spelled words", func(t *testing.T) {
		expected := 76
		got := CalibrationValue("7pqrstsixteen", true)
		assertEqual(t, expected, got)
	})

	t.Run("oneight with filtering spelled words", func(t *testing.T) {
		expected := 18
		got := CalibrationValue("oneight", true)
		assertEqual(t, expected, got)
	})

	t.Run("one with filtering spelled words", func(t *testing.T) {
		expected := 11
		got := CalibrationValue("one", true)
		assertEqual(t, expected, got)
	})

	t.Run("two with filtering spelled words", func(t *testing.T) {
		expected := 22
		got := CalibrationValue("two", true)
		assertEqual(t, expected, got)
	})

	t.Run("three with filtering spelled words", func(t *testing.T) {
		expected := 33
		got := CalibrationValue("three", true)
		assertEqual(t, expected, got)
	})

	t.Run("four with filtering spelled words", func(t *testing.T) {
		expected := 44
		got := CalibrationValue("four", true)
		assertEqual(t, expected, got)
	})

	t.Run("five with filtering spelled words", func(t *testing.T) {
		expected := 55
		got := CalibrationValue("five", true)
		assertEqual(t, expected, got)
	})

	t.Run("six with filtering spelled words", func(t *testing.T) {
		expected := 66
		got := CalibrationValue("six", true)
		assertEqual(t, expected, got)
	})
	t.Run("seven with filtering spelled words", func(t *testing.T) {
		expected := 77
		got := CalibrationValue("seven", true)
		assertEqual(t, expected, got)
	})
}

func TestTotalCalibrationValue(t *testing.T) {
	t.Run("example without filtering spelled words", func(t *testing.T) {
		input := `1abc2
		pqr3stu8vwx
		a1b2c3d4e5f
		treb7uchet`

		expected := 142
		got := TotalCalibrationValue(input, false)

		assertEqual(t, expected, got)
	})

	t.Run("example with filtering spelled words", func(t *testing.T) {
		input := `two1nine
		eightwothree
		abcone2threexyz
		xtwone3four
		4nineeightseven2
		zoneight234
		7pqrstsixteen`

		expected := 281
		got := TotalCalibrationValue(input, true)

		assertEqual(t, expected, got)
	})
}

func TestProcessLineForSpelledNumbers(t *testing.T) {
	t.Run("twone", func(t *testing.T) {
		expected := "21"
		got := ProcessLineForSpelledNumbers("twone")

		assertEqualStr(t, expected, got)
	})

	t.Run("sevenine", func(t *testing.T) {
		expected := "79"
		got := ProcessLineForSpelledNumbers("sevenine")

		assertEqualStr(t, expected, got)
	})

	t.Run("onethree", func(t *testing.T) {
		expected := "13"
		got := ProcessLineForSpelledNumbers("onethree")

		assertEqualStr(t, expected, got)
	})

	t.Run("two1nine", func(t *testing.T) {
		expected := "219"
		got := ProcessLineForSpelledNumbers("two1nine")

		assertEqualStr(t, expected, got)
	})

	t.Run("eightwothree", func(t *testing.T) {
		expected := "823"
		got := ProcessLineForSpelledNumbers("eightwothree")

		assertEqualStr(t, expected, got)
	})

	t.Run("abcone2threexyz", func(t *testing.T) {
		expected := "123"
		got := ProcessLineForSpelledNumbers("abcone2threexyz")

		assertEqualStr(t, expected, got)
	})

	t.Run("xtwone3four", func(t *testing.T) {
		expected := "2134"
		got := ProcessLineForSpelledNumbers("xtwone3four")

		assertEqualStr(t, expected, got)
	})
}
