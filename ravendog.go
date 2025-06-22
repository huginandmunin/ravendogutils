// Package ravendogutils has utility fucntions for dog.
package ravendogutils

import (
	"fmt"
)

// Years calculates dog years based on human years input.
// Ouput is an int.
func Years(humanYears int) int (
	return humanYears * 7
)
