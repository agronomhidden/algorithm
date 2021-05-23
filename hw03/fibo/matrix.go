package main

import (
	"math/big"
	"strconv"
)



var initMatrix = &matrix2d{
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(0),
}

type matrix2d struct {
	a *big.Int
	b *big.Int
	c *big.Int
	d *big.Int
}

func matrixMul(x, y *matrix2d) *matrix2d {
	return &matrix2d{
		bi().Add( bi().Mul(x.a, y.a), bi().Mul(x.b, y.c)),
		bi().Add( bi().Mul(x.a, y.b), bi().Mul(x.b, y.d)),

		bi().Add( bi().Mul(x.c, y.a), bi().Mul(x.d, y.c)),
		bi().Add( bi().Mul(x.c, y.b), bi().Mul(x.d, y.d)),
	}
}

type MatrixCounter struct {

}
func (e *MatrixCounter) Run(input []string) string {
	n, _ := strconv.Atoi(input[0])
	if n <= 1 {
		return big.NewInt(int64(n)).String()
	}
	var result *matrix2d = nil
	base := n - 1
	counter := initMatrix
	i := 2
	for {
		if base%2 != 0 {
			if result == nil {
				result = counter
			} else {
				result = matrixMul(result, counter)
			}
		}
		if base == 1 {
			break
		}
		counter = matrixMul(counter, counter)
		i<<=2
		base = int(base/2)
	}

	return result.a.String()
}
