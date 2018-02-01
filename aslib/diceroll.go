package aslib

import (
	"fmt"
	"math/rand"
)

func Diceroll() (string, error) {

	d1 := rand.Intn(6) + 1
	d2 := rand.Intn(6) + 1

	return fmt.Sprintf("(%d + %d) = %d", d1, d2, d1+d2), nil
}
