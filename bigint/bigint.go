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
