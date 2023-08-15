package repository

//go:generate mockgen -destination=mock/mock_repository.go -package=mock . InterGalacticConverter
type InterGalacticConverter interface {
	DefineUnit(unit string, roman string) error
	DefineMaterial(units []string, material string, credits int)
	ConvertToIntergalactic(inputStr string) string
	ConvertToCredits(unitStr string, material string) (float32, error)
	ConvertRomanToArabic(roman string) int
}

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
