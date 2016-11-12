package main

import (
	"fmt"
	"math/rand"
)

func diceroll() (string, error) {

	d1 := rand.Intn(5) + 1
	d2 := rand.Intn(5) + 1

	return fmt.Sprintf("(%d + %d) = %d", d1, d2, d1+d2), nil
}

