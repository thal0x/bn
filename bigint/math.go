package bigint

import "math/big"

func Add(x, y *big.Int) *big.Int {
	return big.NewInt(0).Add(x, y)
}

func Sub(x, y *big.Int) *big.Int {
	return big.NewInt(0).Sub(x, y)
}

func Mul(x, y *big.Int) *big.Int {
	return big.NewInt(0).Mul(x, y)
}

func DivDown(x, y *big.Int) *big.Int {
	return big.NewInt(0).Quo(x, y)
}

func DivUp(x, y *big.Int) *big.Int {
	return Add(ONE, DivDown(Sub(x, ONE), y))
}
