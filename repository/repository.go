package repository

import (
	"fmt"
	"strings"
)

func (ic *IntergalacticConverter) DefineUnit(unit, roman string) {
	ic.unitToRoman[unit] = roman
}

func (ic *IntergalacticConverter) DefineMaterial(units []string, material string, credits int) {
	roman := ""
	for _, unit := range units {
		roman += ic.unitToRoman[unit]
	}

	ic.materials[material] = credits / ic.ConvertRomanToArabic(roman)
}

func (ic *IntergalacticConverter) ConvertToIntergalactic(inputStr string) string {
	units := strings.Split(inputStr, " ")
	roman := ""
	for _, unit := range units {
		roman += ic.unitToRoman[unit]
	}
	return roman
}

func (ic *IntergalacticConverter) ConvertToCredits(unitStr string, material string) (float32, error) {
	units := strings.Split(unitStr, " ")
	roman := ""
	for _, unit := range units {
		roman += ic.unitToRoman[unit]
	}

	if credits, ok := ic.materials[material]; ok {
		return float32(ic.ConvertRomanToArabic(roman)) * float32(credits), nil
	}
	return 0, fmt.Errorf("material not found")
}

func (ic *IntergalacticConverter) ConvertRomanToArabic(roman string) int {
	romanToArabic := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	arabic := 0
	prevValue := 0

	for i := len(roman) - 1; i >= 0; i-- {
		symbol := string(roman[i])
		value := romanToArabic[symbol]

		if value >= prevValue {
			arabic += value
		} else {
			arabic -= value
		}

		prevValue = value
	}

	return arabic
}
