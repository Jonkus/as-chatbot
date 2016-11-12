package main

import (
	"fmt"
	"math/rand"
)

func motivehit(args []string) (string) {

	d1 := rand.Intn(5) + 1
	d2 := rand.Intn(5) + 1

	roll := d1 + d2
	var movementMode string
	if len(args) < 1 {
	    movementMode = "t"
	} else {
	    movementMode = args[0]
	}

	mod := 0
	switch movementMode {
	case "t", "s", "n":
		mod = 0
	case "w", "h":
		mod = 1
	case "v", "g":
		mod = 2
	default:
		return("unknown movement Mode " + movementMode)
	}

	effect := "No Effect"

	if (roll + mod) > 12 {
		effect = "Unit Immobilized"
	} else if (roll + mod) > 10 {
		effect = "minus 50% Move (round down)"
	} else if (roll + mod) > 8 {
		effect = "minus 2 inch Move (round down)"
	}

	return fmt.Sprintf("[ (%d + %d) + %d (%s) = %d ]\n--->%s<----", d1, d2, mod, movementMode, roll+mod, effect)
}
