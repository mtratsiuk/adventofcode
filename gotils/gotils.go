package gotils

import (
	"fmt"
	"strconv"
)

func MustParseInt(s string) int {
	r, err := strconv.Atoi(s)

	if err != nil {
		panic(fmt.Sprintf("ParseInt: %v", err))
	}

	return r
}
