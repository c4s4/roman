package main

import (
	"fmt"
	"os"
	"strconv"
)

// Mappings for roman figures
var Mappings = map[int]map[string]string{
	// units
	0: {
		"1": "I",
		"2": "II",
		"3": "III",
		"4": "IV",
		"5": "V",
		"6": "VI",
		"7": "VII",
		"8": "VIII",
		"9": "IX",
		"0": "",
	},
	// decades
	1: {
		"1": "X",
		"2": "XX",
		"3": "XXX",
		"4": "XL",
		"5": "L",
		"6": "LX",
		"7": "LXX",
		"8": "LXXX",
		"9": "XC",
		"0": "",
	},
	// hundreds
	2: {
		"1": "C",
		"2": "CC",
		"3": "CCC",
		"4": "CD",
		"5": "D",
		"6": "DC",
		"7": "DCC",
		"8": "DCCC",
		"9": "CM",
		"0": "",
	},
	// thousands
	3: {
		"1": "M",
		"2": "MM",
		"3": "MMM",
		"4": "MMMM",
	},
}

// Max is the maximum roman number
var Max = 5000

// Roman converts given integer to roman
// - arab: is the number to convert in arabian form
// Return: roman number and error if any
func Roman(arab string) (string, error) {
	i, err := strconv.Atoi(arab)
	if err != nil {
		return "", fmt.Errorf("%s is not a valid arabian number", arab)
	}
	if i <= 0 {
		return "", fmt.Errorf("Romans can't count above 1")
	}
	if i >= Max {
		return "", fmt.Errorf("Romans can't count beyond 4999")
	}
	str := strconv.Itoa(i)
	index := 0
	roman := ""
	for str != "" {
		last := str[len(str)-1:]
		str = str[:len(str)-1]
		roman = Mappings[index][last] + roman
		index++
	}
	return roman, nil
}

// Arab converts from roman to arab
// - roman: roman number to convert
// Return arab number and an error if any
func Arab(roman string) (string, error) {
	for i := 1; i < Max; i++ {
		candidate, _ := Roman(strconv.Itoa(i))
		if candidate == roman {
			return strconv.Itoa(i), nil
		}
	}
	return "", fmt.Errorf("%s is not a valid roman number", roman)
}

func main() {
	for _, arab := range os.Args[1:] {
		roman, err := Roman(arab)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s -> %s\n", arab, roman)
	}
}
