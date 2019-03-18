package main

import (
	"testing"
)

var Numbers = map[string]string{
	"1":    "I",
	"2":    "II",
	"3":    "III",
	"4":    "IV",
	"5":    "V",
	"6":    "VI",
	"7":    "VII",
	"8":    "VIII",
	"9":    "IX",
	"10":   "X",
	"11":   "XI",
	"12":   "XII",
	"45":   "XLV",
	"50":   "L",
	"100":  "C",
	"500":  "D",
	"1000": "M",
	"1999": "MCMXCIX",
	"3457": "MMMCDLVII",
}

func TestRoman(t *testing.T) {
	for number, expected := range Numbers {
		actual, _ := Roman(number)
		if actual != expected {
			t.Errorf("Roman(%s) should be '%s', '%s' instead\n", number, expected, actual)
		}
	}
}

func TestRomanErrors(t *testing.T) {
	_, err := Roman("-1")
	if err.Error() != "Romans can't count above 1" {
		t.Errorf("Bad error message: %s", err.Error())
	}
	_, err = Roman("0")
	if err.Error() != "Romans can't count above 1" {
		t.Errorf("Bad error message: %s", err.Error())
	}
	_, err = Roman("5000")
	if err.Error() != "Romans can't count beyond 4999" {
		t.Errorf("Bad error message: %s", err.Error())
	}
}

func TestArab(t *testing.T) {
	for expected, roman := range Numbers {
		actual, _ := Arabic(roman)
		if actual != expected {
			t.Errorf("Arab(%s) should be '%s', '%s' instead\n", roman, expected, actual)
		}
	}
}

func TestArabErrors(t *testing.T) {
	_, err := Arabic("Mickey")
	if err.Error() != "Mickey is not a valid roman number" {
		t.Errorf("Bad error message: %s", err.Error())
	}
}
