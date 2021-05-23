package main

import (
	"fmt"
	"github.com/agronomhidden/algorithm/tester"

	"strconv"
	"strings"

)

type PowerCounter struct {

}
func (e *PowerCounter) Run(input []string) string {
	n, _ := strconv.ParseFloat(input[0],64)
	p, _ := strconv.Atoi(input[1])
	if p == 0 {
		return "1.0"
	}
	result := n
	for i := 2; i <= p; i++ {
		result *= n
	}
	return formatFloatResult(result)
}

type MultiplicationCounter struct {

}
func (e *MultiplicationCounter) Run(input []string) string {
	n, _ := strconv.ParseFloat(input[0],64)
	p, _ := strconv.Atoi(input[1])
	if p == 0 {
		return "1.0"
	}
	result := n
	i := 2
	for j := i; j <= p; j <<=1{
		result *= result
		i = j
	}
	for k := i; k < p; k++ {
		result *= n
	}
	return formatFloatResult(result)
}

type BinaryCounter struct {

}
func (e *BinaryCounter) Run(input []string) string {
	n, _ := strconv.ParseFloat(input[0],64)
	p, _ := strconv.Atoi(input[1])
	if p == 0 {
		return "1.0"
	}
	base := p
	result := 1.0
	counter := n
	i := 2
	for {
		if base%2 != 0 {
			result *= counter
		}
		if base == 1 {
			break
		}
		counter *= counter
		i<<=2
		base = int(base/2)
	}

	return formatFloatResult(result)
}

func formatFloatResult(r float64) string {
	str := strconv.FormatFloat(r, 'f', 11, 64)
	str = strings.TrimRight(str, "0")
	if r == float64(int(r)) {
		str += "0"
	}
	return str
}
func main()  {
	fmt.Println("Binary counter:")
	tester.Run("./hw03/power/testdata", &BinaryCounter{}, tester.Option{})
	fmt.Println("Multiplication counter:")
	tester.Run("./hw03/power/testdata", &MultiplicationCounter{}, tester.Option{})
	fmt.Println("Iteration counter:")
	tester.Run("./hw03/power/testdata", &PowerCounter{}, tester.Option{})
}