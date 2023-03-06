package u256

import "github.com/holiman/uint256"

func IsEqual(x, y *uint256.Int) bool {
	return x.Cmp(y) == 0
}

func IsNotEqual(x, y *uint256.Int) bool {
	return x.Cmp(y) != 0
}

func IsLessThan(x, y *uint256.Int) bool {
	return x.Cmp(y) < 0
}

func IsLessThanOrEqual(x, y *uint256.Int) bool {
	return x.Cmp(y) <= 0
}

func IsGreaterThan(x, y *uint256.Int) bool {
	return x.Cmp(y) > 0
}

func IsGreaterThanOrEqual(x, y *uint256.Int) bool {
	return x.Cmp(y) >= 0
}
