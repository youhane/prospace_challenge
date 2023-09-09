package helpers

import (
	"fmt"
	"math"
	"strconv"
)

func FormatCredits(credits float32) string {
	if math.Mod(float64(credits), 1.0) < 0.0001 {
		return strconv.Itoa(int(credits))
	} else {
		return fmt.Sprintf("%.2f", credits)
	}
}
