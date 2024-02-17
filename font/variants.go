package font

import "github.com/dlclark/regexp2"

const (
	variantRx string = `^([1-9]00[iI]?(?!.))|regular|italic|all$`
	variantGroupRx string = `^[\w,]*$`
)

/**
 * Validates the actual "variants" argument
*/
func ValidateVariantsArg(vars string) bool {
	re := regexp2.MustCompile(variantGroupRx, 0)
	isMatched, _ := re.MatchString(vars)
	return isMatched
}

/**
 * Validates variants one by one, if one or more is invalid returns the indices of these variants in an array
*/
func ValidateVariants(vars []string) []int {
	var indices []int

	for i, variant := range vars {
		re :=	regexp2.MustCompile(variantRx, 0)

		matched, _ := re.MatchString(variant)

		if !matched {
			indices = append(indices, i)
		}
	}

	if len(indices) == 0 {
		return nil
	}
	return indices
}
