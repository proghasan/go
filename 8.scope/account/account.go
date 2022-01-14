package account

import "fmt"

var PersonName = "Mohammad Hasan"

func PrintMe(mainPackageLevelVariable, mainBlockLevelVariable, name string) {
	fmt.Println(mainBlockLevelVariable, mainBlockLevelVariable, name)
}
