package u256

import (
	"math/big"

	"github.com/holiman/uint256"
)

var (
	ZERO = New(0)
	ONE  = New(1)
)

func New(x uint64) *uint256.Int {
	return uint256.NewInt(x)
}

func FromBig(x *big.Int) *uint256.Int {
	z, _ := uint256.FromBig(x)
	return z
}

func FromString(s string) *uint256.Int {
	x, _ := big.NewInt(0).SetString(s, 0)
	return FromBig(x)
}

func FromBytes(b []byte) *uint256.Int {
	return new(uint256.Int).SetBytes(b)
}
