package bigint

import "math/big"

func Add(x, y *big.Int) *big.Int {
	return new(big.Int).Add(x, y)
}

func Sub(x, y *big.Int) *big.Int {
	return new(big.Int).Sub(x, y)
}

func Mul(x, y *big.Int) *big.Int {
	return new(big.Int).Mul(x, y)
}

func DivDown(x, y *big.Int) *big.Int {
	return new(big.Int).Quo(x, y)
}

func DivUp(x, y *big.Int) *big.Int {
	return Add(ONE, DivDown(Sub(x, ONE), y))
}

func Mod(x, y *big.Int) *big.Int {
	return new(big.Int).Mod(x, y)
}
