package main

import (
	"myapp/account"
)

// package level variable
var mainPackageLevelVariable = "This is main package level variable."

func main() {
	// block level variable
	var mainBlockLevelVariable = "This is main block level variable."

	account.PrintMe(mainPackageLevelVariable, mainBlockLevelVariable, account.PersonName)

}
