package main

import (
	"fmt"

	chubby "github.com/whiteboxsolutions/Chubby"
)

var rolls chubby.Rolls = chubby.New()

func main() {
	adminRoll := rolls.NewRoll("Admin")
	_ = rolls.NewRoll("Manager")
	_ = rolls.NewRoll("User")
	_ = rolls.NewRoll("Anonymous")

	lowSecurityRolls := rolls.Combine("User", "Anonymous")

	seniorRolls := rolls.Combine("Admin", "Manager")

	if chubby.HasRoll(seniorRolls, adminRoll) {
		fmt.Println("adminRoll is in SeniorRolls")
	}

	if !chubby.HasRoll(lowSecurityRolls, adminRoll) {
		fmt.Println("adminRoll is NOT in the lowSecurityRolls")
	}
}
