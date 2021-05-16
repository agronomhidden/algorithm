package main

import (
	"github.com/agronomhidden/algorithm/tester"
	"strconv"
)

type LengthCounter struct {

}
func (e *LengthCounter) Run(input []string) string {
	return strconv.Itoa(len(input[0]))
}

func main()  {
	tester.Run("./hw02/strings/testdata", &LengthCounter{})
}

