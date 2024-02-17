package font

import (
	"fmt"
	"os"
	"strings"
)

func PrintVariantsErrAndExit(variants []string, indices []int) {
	fmt.Println("Following variants are invalid: ")

	fmt.Print("[")

	var str strings.Builder

	for i, index := range indices {

		str.WriteString(variants[index])

		if i < len(indices) - 1 {
			str.WriteString(", ")
		}		
	}

	fmt.Print(str.String())

	fmt.Print("]")

	os.Exit(1)
}
