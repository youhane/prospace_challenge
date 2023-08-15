package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type IntergalacticConverter struct {
	unitToRoman map[string]string
	materials   map[string]int
}

func NewIntergalacticConverter() *IntergalacticConverter {
	return &IntergalacticConverter{
		unitToRoman: make(map[string]string),
		materials:   make(map[string]int),
	}
}

func (ic *IntergalacticConverter) DefineUnit(unit, roman string) {
	ic.unitToRoman[unit] = roman
}

func (ic *IntergalacticConverter) DefineMaterial(units []string, material string, credits int) {
	roman := ""
	for _, unit := range units {
		roman += ic.unitToRoman[unit]
	}

	ic.materials[material] = credits / ic.convertRomanToArabic(roman)
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
		return float32(ic.convertRomanToArabic(roman)) * float32(credits), nil
	}
	return 0, fmt.Errorf("material not found")
}

func (ic *IntergalacticConverter) convertRomanToArabic(roman string) int {
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

func main() {
	converter := NewIntergalacticConverter()

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter unit definitions (one per line, blank line to stop):")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		parts := strings.Fields(input)
		if len(parts) != 3 || parts[1] != "is" {
			fmt.Println("Invalid input. Expected format: <unit> is <roman>")
			continue
		}
		converter.DefineUnit(parts[0], parts[2])
	}

	fmt.Println("Enter credit definitions (one per line, blank line to stop):")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		parts := strings.Fields(input)
		if len(parts) < 5 || parts[len(parts)-3] != "is" || parts[len(parts)-1] != "Credits" {
			fmt.Println("Invalid input. Expected format: <unit> ... <unit> <material> is <credits> Credits")
			continue
		}
		material := parts[len(parts)-4]

		var unitStr []string
		for i := 0; i < len(parts); i++ {
			if parts[i] == material {
				break
			}
			unitStr = append(unitStr, parts[i])
		}

		credits := 0
		fmt.Sscanf(parts[len(parts)-2], "%d", &credits)
		converter.DefineMaterial(unitStr, material, credits)
	}

	fmt.Println("Enter questions (one per line, blank line to stop):")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "" {
			break
		}
		parts := strings.Fields(input)

		if parts[0] == "how" && parts[1] == "much" && parts[2] == "is" {
			// how much is ...
			intergalacticUnit := strings.Join(parts[3:], " ")
			intergalacticUnit = strings.TrimRight(intergalacticUnit, " ?")

			romanNumber := converter.ConvertToIntergalactic(intergalacticUnit)
			arabicNumber := converter.convertRomanToArabic(romanNumber)

			fmt.Printf("%s is %d\n", intergalacticUnit, arabicNumber)
		} else if parts[0] == "how" && parts[1] == "many" {
			// how many Credits is ...
			intergalacticUnit := strings.Join(parts[4:len(parts)-2], " ")
			material := parts[len(parts)-2]

			credits, err := converter.ConvertToCredits(intergalacticUnit, material)

			if err != nil {
				fmt.Printf("Error: %s\n", err)
			} else {
				fmt.Printf("%s %s is %.2f Credits\n", intergalacticUnit, material, credits)
			}
		} else if parts[0] == "Does" && strings.Contains(input, "has more Credits") || strings.Contains(input, "have more Credits") || strings.Contains(input, "has less Credits") || strings.Contains(input, "have less Credits") {
			// Does ... have more/less Credits than ...
			// Find the index of the word "has" or "have"
			hasIndex := -1
			for i := 0; i < len(parts); i++ {
				if parts[i] == "has" || parts[i] == "have" {
					hasIndex = i
					break
				}
			}

			firstUnit := strings.Join(parts[1:hasIndex], " ")

			// Find the index for the word "Credits"
			creditsIndex := -1
			for i := 0; i < len(parts); i++ {
				if parts[i] == "Credits" {
					creditsIndex = i
					break
				}
			}

			secondUnit := strings.Join(parts[creditsIndex+1:], " ")
			secondUnit = strings.TrimRight(secondUnit, " ?")

			materialFirst := parts[hasIndex-1]
			materialSecond := parts[len(parts)-2]

			creditsFirst, _ := converter.ConvertToCredits(firstUnit, materialFirst)
			creditsSecond, _ := converter.ConvertToCredits(secondUnit, materialSecond)

			if (strings.Contains(input, "has more Credits") || strings.Contains(input, "have more Credits")) && creditsFirst > creditsSecond {
				fmt.Printf("%s has more Credits than %s\n", firstUnit, secondUnit)
			} else if (strings.Contains(input, "has less Credits") || strings.Contains(input, "have less Credits")) && creditsFirst < creditsSecond {
				fmt.Printf("%s has less Credits than %s\n", firstUnit, secondUnit)
			} else {
				fmt.Printf("%s has equal Credits as %s\n", firstUnit, secondUnit)
			}
		} else if parts[0] == "Is" && (strings.Contains(input, "larger than") || strings.Contains(input, "smaller than")) {
			// Is ... larger/smaller than ...
			// Find the index for the word "than"
			thanIndex := -1
			for i := 0; i < len(parts); i++ {
				if parts[i] == "than" {
					thanIndex = i
					break
				}
			}

			firstUnit := strings.Join(parts[1:thanIndex-1], " ")
			secondUnit := strings.Join(parts[thanIndex+1:len(parts)-1], " ")

			creditsFirst := converter.ConvertToIntergalactic(firstUnit)
			creditsSecond := converter.ConvertToIntergalactic(secondUnit)

			var comparison string
			if creditsFirst > creditsSecond {
				comparison = "larger than"
			} else if creditsFirst < creditsSecond {
				comparison = "smaller than"
			} else {
				comparison = "equal to"
			}
			fmt.Printf("%s is %s %s\n", firstUnit, comparison, secondUnit)
		} else {
			fmt.Println("I have no idea what you are talking about")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
