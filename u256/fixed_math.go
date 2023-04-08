package u256

import "github.com/holiman/uint256"

func FixedMulUp(x, y *uint256.Int) *uint256.Int {
	product := Mul(x, y)

	if IsEqual(product, ZERO) {
		return product
	}

	return Add(
		DivDown(
			Sub(product, ONE),
			ONE_18,
		),
		ONE,
	)
}

func FixedMulDown(x, y *uint256.Int) *uint256.Int {
	return DivDown(Mul(x, y), ONE_18)
}

func FixedDivDown(x, y *uint256.Int) *uint256.Int {
	if IsEqual(x, ZERO) {
		return x
	}

	return DivDown(Mul(x, ONE_18), y)
}

func FixedDivUp(x, y *uint256.Int) *uint256.Int {
	if IsEqual(x, ZERO) {
		return x
	}

	return Add(
		DivDown(
			Sub(
				Mul(x, ONE_18),
				ONE,
			),
			y,
		),
		ONE,
	)
}

func FixedPowUp(x, y *uint256.Int) *uint256.Int {
	// if IsEqual(y, ONE_18) {
	// 	return x
	// }

	// if IsEqual(y, TWO_18) {
	// 	return FixedMulUp(x, x)
	// }

	// if IsEqual(y, FOUR_18) {
	// 	square := FixedMulUp(x, x)
	// 	return FixedMulUp(square, square)
	// }

	raw := Pow(x, y)
	maxError := Add(FixedMulUp(raw, uint256.NewInt(10000)), ONE)
	return Add(raw, maxError)
}
