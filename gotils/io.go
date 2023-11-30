package gotils

import (
	"os"
	"path"
	"strings"
)

func ReadInput(name string) string {
	data := os.Getenv("AOC_DATA_PATH")

	input, err := os.ReadFile(path.Join(data, name+".txt"))

	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(string(input))
}
