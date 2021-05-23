package main

import (
	"flag"
	"fmt"
	"github.com/agronomhidden/algorithm/tester"
	"math"
	"math/big"
	"strconv"
)
func bi() *big.Int {
	return big.NewInt(0)
}
func bf() *big.Float {
	return big.NewFloat(0)
}

type RecursiveCounter struct {

}
func (e *RecursiveCounter) Run(input []string) string {
	n, _ := strconv.Atoi(input[0])

	return fiboRecursive(n).String()
}
func fiboRecursive(n int) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}
	return bi().Add(fiboRecursive(n - 1), fiboRecursive(n - 2))
}

type IterationCounter struct {

}
func (e *IterationCounter) Run(input []string) string {
	n, _ := strconv.Atoi(input[0])
	F0 := big.NewInt(0)
	F1 := big.NewInt(1)
	F2 := big.NewInt(0)
	if n <= 1 {
		return big.NewInt(int64(n)).String()
	}
	for i := 2; i <= n; i++ {
		F2 = bi().Add(F1, F0)
		F0 = F1
		F1 = F2
	}
	return F2.String()
}

type FormulaCounter struct {

}
func (e *FormulaCounter) Run(input []string) string {
	n, _ := strconv.Atoi(input[0])
	phi := big.NewFloat(math.Phi)
	add := big.NewFloat(0.5)
	del := big.NewFloat(math.Sqrt(5))

	result, _ := bf().Add(
			bf().Quo(
				bigPow(phi, n), del,
			),
			add,
		).Int(bi())
	return result.String()
}
func bigPow(dig *big.Float, p int) *big.Float {
	if p == 0 {
		return big.NewFloat(1)
	}
	base := p
	result := big.NewFloat(1)
	counter := dig
	i := 2
	for {
		if base%2 != 0 {
			result = bf().Mul(result, counter)
		}
		if base == 1 {
			break
		}
		counter = bf().Mul(counter, counter)
		i<<=2
		base = int(base/2)
	}
	return result
}

var mode string
const (
	r = "r"
	i = "i"
	f = "f"
	m = "m"
)
func init() {
	flag.StringVar(&mode, "m", m, "running different algorithms")
	flag.Parse()
}

func main()  {
	switch mode {
	case r:
		fmt.Println("Recursive counter:")
		tester.Run("./hw03/fibo/testdata", &RecursiveCounter{}, tester.Option{})
	case i:
		fmt.Println("Iteration counter:")
		tester.Run("./hw03/fibo/testdata", &IterationCounter{}, tester.Option{})
	case f:
		fmt.Println("Formula counter:")
		tester.Run("./hw03/fibo/testdata", &FormulaCounter{}, tester.Option{PreventOutput: true})
	case m:
		fmt.Println("Matrix counter:")
		tester.Run("./hw03/fibo/testdata", &MatrixCounter{}, tester.Option{})
	}
}
