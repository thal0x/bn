package u256

import (
	"errors"

	"github.com/holiman/uint256"
)

func Add(x, y *uint256.Int) *uint256.Int {
	return new(uint256.Int).Add(x, y)
}

func SafeAdd(x, y *uint256.Int) (*uint256.Int, error) {
	result, overflow := new(uint256.Int).AddOverflow(x, y)
	if overflow {
		return nil, errors.New("SafeAdd: overflow")
	}

	return result, nil
}

func Sub(x, y *uint256.Int) *uint256.Int {
	return new(uint256.Int).Sub(x, y)
}

func SafeSub(x, y *uint256.Int) (*uint256.Int, error) {
	result, overflow := new(uint256.Int).SubOverflow(x, y)
	if overflow {
		return nil, errors.New("SafeSub: overflow")
	}

	return result, nil
}

func Mul(x, y *uint256.Int) *uint256.Int {
	return new(uint256.Int).Mul(x, y)
}

func SafeMul(x, y *uint256.Int) (*uint256.Int, error) {
	result, overflow := new(uint256.Int).MulOverflow(x, y)
	if overflow {
		return nil, errors.New("SafeMul: overflow")
	}

	return result, nil
}

func Div(x, y *uint256.Int, roundUp bool) *uint256.Int {
	if roundUp {
		return DivUp(x, y)
	}

	return DivDown(x, y)
}

func DivDown(x, y *uint256.Int) *uint256.Int {
	return new(uint256.Int).Div(x, y)
}

func DivUp(x, y *uint256.Int) *uint256.Int {
	return Add(ONE, DivDown(Sub(x, ONE), y))
}

func Mod(x, y *uint256.Int) *uint256.Int {
	return new(uint256.Int).Mod(x, y)
}

func MulMod(x, y, m *uint256.Int) *uint256.Int {
	return new(uint256.Int).MulMod(x, y, m)
}

func MulDiv(a, b, denominator *uint256.Int) *uint256.Int {
	result, _ := new(uint256.Int).MulDivOverflow(a, b, denominator)

	return result
}

func MulDivRoundingUp(a, b, denominator *uint256.Int) *uint256.Int {
	result := MulDiv(a, b, denominator)

	if IsGreaterThan(MulMod(a, b, denominator), ZERO) {
		result = Add(result, ONE)
	}

	return result
}

func Rsh(x *uint256.Int, n uint) *uint256.Int {
	return new(uint256.Int).Rsh(x, n)
}

func Lsh(x *uint256.Int, n uint) *uint256.Int {
	return new(uint256.Int).Lsh(x, n)
}

func And(x, y *uint256.Int) *uint256.Int {
	return new(uint256.Int).And(x, y)
}

func Or(x, y *uint256.Int) *uint256.Int {
	return new(uint256.Int).Or(x, y)
}

func Xor(x, y *uint256.Int) *uint256.Int {
	return new(uint256.Int).Xor(x, y)
}

func Not(x *uint256.Int) *uint256.Int {
	return new(uint256.Int).Not(x)
}
