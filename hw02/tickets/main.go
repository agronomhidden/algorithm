package main

import (
	"fmt"
	"github.com/agronomhidden/algorithm/tester"
	"strconv"
)
const BASE = 10

type RecursiveCounter struct {

}

func (e *RecursiveCounter) Run(input []string) string {
	n, _ := strconv.Atoi(input[0])
	sum := counter{n, 0}
	sum.nextDigit(0,0,0)

	return strconv.Itoa(sum.counter)
}
type counter struct {
	n int
	counter int
}

func (e *counter) nextDigit(digit, sum1, sum2 int) {
	if digit == e.n {
		if sum1 == sum2 {
			e.counter++
		}
		return
	}
	for a := 0; a < BASE; a++ {
		for b := 0; b < BASE; b++ {
			e.nextDigit(digit + 1, sum1 + a, sum2 + b)
		}
	}
}

type Dinamic struct {

}

func (e *Dinamic) Run(input []string) string {
	n, _ := strconv.Atoi(input[0])
	if n < 1 {
		return "0"
	}
	if n == 1 {
		return strconv.Itoa(BASE)
	}
	seq := initSequence()
	var matrixInst *matrix
	for x := 2; x <= n; x++ {
		matrixInst = matrixNew(n, seq)
		seq = matrixInst.getColumnsSum()
	}

	return strconv.FormatInt(matrixInst.getSum(), 10)
}

func matrixNew(n int, sequence []int64) *matrix {
	w := n*BASE - n + 1
	inst := matrix{
		width: w,
		data: make([][]int64, BASE, BASE),
	}
	for i := 0; i < BASE; i++ {
		inst.data[i] = make([]int64, w, w)
		for k := 0; k < w; k++ {
			if k < len(sequence) && k+i < w {
				inst.data[i][k+i] = sequence[k]
			}
		}
	}
	return &inst
}

type matrix struct {
	width int
	data [][]int64
}

func (e matrix) getColumnsSum() (seq []int64) {
	for k := 0; k < e.width; k++ {
		var sum int64
		for i := 0; i < BASE; i++ {
			sum += e.data[i][k]
		}
		seq = append(seq, sum)
	}
	return
}

func (e matrix) getSum() (result int64) {
	for _, value := range e.getColumnsSum() {
		result += value*value
	}
	return
}

func initSequence() (seq []int64) {
	for i := 0; i < BASE; i++ {
		seq = append(seq, 1)
	}
	return seq
}

func main()  {
	fmt.Println("Dinamic counter:")
	tester.Run("./hw02/tickets/testdata", &Dinamic{}, tester.Option{})
	fmt.Println("Recursive counter (wait if you want...):")
	tester.Run("./hw02/tickets/testdata", &RecursiveCounter{}, tester.Option{})
}

