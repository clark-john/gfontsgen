package font

import (
	"fmt"
	"os"
	"strings"
)

func _PrintVariantsErr(
	variants []string, 
	indices []int, 
	isExit bool,
	printMessage string,
){
	fmt.Print(printMessage)

	fmt.Print("[")

	var str strings.Builder

	for i, index := range indices {

		str.WriteString(variants[index])

		if i < len(indices) - 1 {
			str.WriteString(", ")
		}		
	}

	fmt.Print(str.String())

	fmt.Println("]")

	if isExit {
		os.Exit(1)
	}	
}

func PrintVariantsErrWithName(variants []string, indices []int, isExit bool, family string){
	_PrintVariantsErr(
		variants, indices, isExit, 
		fmt.Sprintf("Following variants are invalid in font %s: ", family),
	)
}

func PrintVariantsErr(variants []string, indices []int, isExit bool) {
	_PrintVariantsErr(variants, indices, isExit, "Following variants are invalid: ")
}
