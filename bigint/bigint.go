package bigint

import (
	"math/big"
)

var (
	ZERO = New(0)
	ONE  = New(1)

	ONE_18 = FromString("1000000000000000000") // 1e18
)

func New(x int64) *big.Int {
	return big.NewInt(x)
}

func FromString(s string) *big.Int {
	x, _ := big.NewInt(0).SetString(s, 0)
	return x
}

func FromBytes(b []byte) *big.Int {
	return big.NewInt(0).SetBytes(b)
}

func Neg(x *big.Int) *big.Int {
	return new(big.Int).Neg(x)
}

func Rsh(x *big.Int, n uint) *big.Int {
	return new(big.Int).Rsh(x, n)
}

func Lsh(x *big.Int, n uint) *big.Int {
	return new(big.Int).Lsh(x, n)
}

func And(x, y *big.Int) *big.Int {
	return new(big.Int).And(x, y)
}

func Or(x, y *big.Int) *big.Int {
	return new(big.Int).Or(x, y)
}

func Xor(x, y *big.Int) *big.Int {
	return new(big.Int).Xor(x, y)
}

func AndNot(x, y *big.Int) *big.Int {
	return new(big.Int).AndNot(x, y)
}

func Not(x *big.Int) *big.Int {
	return new(big.Int).Not(x)
}
