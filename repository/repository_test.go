package repository

import (
	"strings"
	"testing"
)

func TestNewIntergalacticConverter(t *testing.T) {
	converter := NewIntergalacticConverter()

	if converter == nil {
		t.Errorf("NewIntergalacticConverter should not return nil")
	}
}

func TestDefineUnit(t *testing.T) {
	converter := &IntergalacticConverter{
		unitToRoman: make(map[string]string),
	}

	converter.DefineUnit("glob", "I")

	expected := "I"
	if converter.unitToRoman["glob"] != expected {
		t.Errorf("Expected unit 'glob' to be defined as %s, but got %s", expected, converter.unitToRoman["glob"])
	}
}

func TestDefineMaterial(t *testing.T) {
	converter := &IntergalacticConverter{
		unitToRoman: make(map[string]string),
		materials:   make(map[string]int),
	}

	converter.DefineUnit("glob", "I")
	converter.DefineUnit("prok", "V")

	converter.DefineMaterial([]string{"glob", "glob"}, "Silver", 17)

	expected := 17 / converter.ConvertRomanToArabic("II")
	if converter.materials["Silver"] != expected {
		t.Errorf("Expected material 'Silver' to have %d credits per unit, but got %d", expected, converter.materials["Silver"])
	}
}

func TestConvertToIntergalactic(t *testing.T) {
	converter := &IntergalacticConverter{
		unitToRoman: map[string]string{
			"glob": "I",
			"prok": "V",
			"pish": "X",
			"tegj": "L",
		},
	}

	tests := []struct {
		input    string
		expected string
	}{
		{"glob prok", "IV"},
		{"pish tegj", "XL"},
		{"glob glob glob", "III"},
		{"", ""},
	}

	for _, test := range tests {
		result := converter.ConvertToIntergalactic(test.input)
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %s, Got: %s", test.input, test.expected, result)
		}
	}
}

func TestConvertToCredits(t *testing.T) {
	converter := &IntergalacticConverter{
		unitToRoman: map[string]string{
			"glob": "I",
			"prok": "V",
			"pish": "X",
			"tegj": "L",
		},
		materials: map[string]int{
			"Silver": 17,
			"Gold":   14450,
		},
	}

	tests := []struct {
		unitStr  string
		material string
		expected float32
	}{
		{"glob glob", "Silver", 34},
		{"pish pish", "Gold", 28900},
		{"glob prok", "Gold", 5750},
		{"tegj pish", "Silver", 0},
		{"glob glob glob glob", "Silver", 0}, // Invalid input
		{"pish pish", "Platinum", 0},         // Material not found
	}

	for _, test := range tests {
		result, err := converter.ConvertToCredits(test.unitStr, test.material)
		if err != nil {
			if err.Error() != "material not found" && err.Error() != "requested number is in invalid format" {
				t.Errorf("Unexpected error: %v", err)
			}
		} else {
			roman := ""
			units := strings.Split(test.unitStr, " ")
			for _, unit := range units {
				roman += converter.unitToRoman[unit]
			}
			romanToArabic := converter.ConvertRomanToArabic(roman)
			expected := float32(romanToArabic) * float32(converter.materials[test.material])

			if result != expected {
				t.Errorf("For units %s and material %s, expected %.6f credits, but got %.6f", test.unitStr, test.material, expected, result)
			}
		}
	}
}

func TestConvertRomanToArabic(t *testing.T) {
	converter := &IntergalacticConverter{}

	tests := []struct {
		roman    string
		expected int
	}{
		{"I", 1},
		{"V", 5},
		{"IX", 9},
		{"XXIV", 24},
		{"XLII", 42},
		{"LXXXIX", 89},
		{"CXCIV", 194},
		{"DCCCXC", 890},
		{"CM", 900},
		{"MCMXCIV", 1994},
	}

	for _, test := range tests {
		result := converter.ConvertRomanToArabic(test.roman)
		if result != test.expected {
			t.Errorf("For Roman numeral %s, expected %d, but got %d", test.roman, test.expected, result)
		}
	}
}
