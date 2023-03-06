package u256

import "github.com/holiman/uint256"

func Add(x, y *uint256.Int) *uint256.Int {
	return new(uint256.Int).Add(x, y)
}

func Sub(x, y *uint256.Int) *uint256.Int {
	return new(uint256.Int).Sub(x, y)
}

func Mul(x, y *uint256.Int) *uint256.Int {
	return new(uint256.Int).Mul(x, y)
}

func DivDown(x, y *uint256.Int) *uint256.Int {
	return new(uint256.Int).Div(x, y)
}

func DivUp(x, y *uint256.Int) *uint256.Int {
	return Add(ONE, DivDown(Sub(x, ONE), y))
}
