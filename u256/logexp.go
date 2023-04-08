package u256

import (
	"math/big"

	"github.com/holiman/uint256"

	"github.com/thal0x/bn/bigint"
)

var (
	ONE_20 = bigint.FromString("100000000000000000000")
	ONE_36 = bigint.FromString("1000000000000000000000000000000000000")

	MAX_NATURAL_EXPONENT = bigint.FromString("130000000000000000000")
	MIN_NATURAL_EXPONENT = bigint.FromString("-41000000000000000000")

	LN_36_LOWER_BOUND = bigint.Sub(bigint.ONE_18, bigint.FromString("100000000000000000"))
	LN_36_UPPER_BOUND = bigint.Add(bigint.ONE_18, bigint.FromString("100000000000000000"))

	// 18 decimal constants
	x0 = bigint.FromString("128000000000000000000")
	a0 = bigint.FromString("38877084059945950922200000000000000000000000000000000000")
	x1 = bigint.FromString("64000000000000000000")
	a1 = bigint.FromString("6235149080811616882910000000")

	// 20 decimal constants
	x2  = bigint.FromString("3200000000000000000000")             // 2ˆ5
	a2  = bigint.FromString("7896296018268069516100000000000000") // eˆ(x2)
	x3  = bigint.FromString("1600000000000000000000")             // 2ˆ4
	a3  = bigint.FromString("888611052050787263676000000")        // eˆ(x3)
	x4  = bigint.FromString("800000000000000000000")              // 2ˆ3
	a4  = bigint.FromString("298095798704172827474000")           // eˆ(x4)
	x5  = bigint.FromString("400000000000000000000")              // 2ˆ2
	a5  = bigint.FromString("5459815003314423907810")             // eˆ(x5)
	x6  = bigint.FromString("200000000000000000000")              // 2ˆ1
	a6  = bigint.FromString("738905609893065022723")              // eˆ(x6)
	x7  = bigint.FromString("100000000000000000000")              // 2ˆ0
	a7  = bigint.FromString("271828182845904523536")              // eˆ(x7)
	x8  = bigint.FromString("50000000000000000000")               // 2ˆ(-1)
	a8  = bigint.FromString("164872127070012814685")              // eˆ(x8)
	x9  = bigint.FromString("25000000000000000000")               // 2ˆ(-2)
	a9  = bigint.FromString("128402541668774148407")              // eˆ(x9)
	x10 = bigint.FromString("12500000000000000000")               // 2ˆ(-3)
	a10 = bigint.FromString("113314845306682631683")              // eˆ(x10)
	x11 = bigint.FromString("6250000000000000000")                // 2ˆ(-4)
	a11 = bigint.FromString("106449445891785942956")              // eˆ(x11)
)

func Pow(x, y *uint256.Int) *uint256.Int {
	if IsEqual(y, ZERO) {
		return ONE_18
	}

	if IsEqual(x, ZERO) {
		return ZERO
	}

	x_int256 := x.ToBig()
	y_int256 := y.ToBig()

	var logx_times_y *big.Int

	if bigint.IsLessThan(LN_36_LOWER_BOUND, x_int256) && bigint.IsLessThan(x_int256, LN_36_UPPER_BOUND) {
		ln_36_x := _ln36(x_int256)

		logx_times_y = bigint.Add(
			bigint.Mul(
				bigint.DivDown(
					ln_36_x,
					bigint.ONE_18,
				),
				y_int256,
			),
			bigint.DivDown(
				bigint.Mul(
					bigint.New(ln_36_x.Int64()%bigint.ONE_18.Int64()),
					y_int256,
				),
				bigint.ONE_18,
			),
		)

	} else {
		logx_times_y = bigint.Mul(Ln(x_int256), y_int256)
	}

	logx_times_y = bigint.DivDown(logx_times_y, bigint.ONE_18)

	return FromBig(Exp(logx_times_y))
}

func Exp(x *big.Int) *big.Int {
	if bigint.IsLessThan(x, MIN_NATURAL_EXPONENT) || bigint.IsGreaterThan(x, MAX_NATURAL_EXPONENT) {
		return bigint.ZERO
	}

	if bigint.IsLessThan(x, bigint.ZERO) {
		return bigint.DivDown(bigint.Mul(bigint.ONE_18, bigint.ONE_18), Exp(new(big.Int).Neg(x)))
	}

	var firstAN *big.Int
	if bigint.IsGreaterThanOrEqual(x, x0) {
		x = bigint.Sub(x, x0)
		firstAN = a0
	} else if bigint.IsGreaterThanOrEqual(x, x1) {
		x = bigint.Sub(x, x1)
		firstAN = a1
	} else {
		firstAN = bigint.ONE
	}

	x = bigint.Mul(x, bigint.New(100))
	product := ONE_20

	if bigint.IsGreaterThanOrEqual(x, x2) {
		x = bigint.Sub(x, x2)
		product = bigint.DivDown(bigint.Mul(product, a2), ONE_20)
	}

	if bigint.IsGreaterThanOrEqual(x, x3) {
		x = bigint.Sub(x, x3)
		product = bigint.DivDown(bigint.Mul(product, a3), ONE_20)
	}

	if bigint.IsGreaterThanOrEqual(x, x4) {
		x = bigint.Sub(x, x4)
		product = bigint.DivDown(bigint.Mul(product, a4), ONE_20)
	}

	if bigint.IsGreaterThanOrEqual(x, x5) {
		x = bigint.Sub(x, x5)
		product = bigint.DivDown(bigint.Mul(product, a5), ONE_20)
	}

	if bigint.IsGreaterThanOrEqual(x, x6) {
		x = bigint.Sub(x, x6)
		product = bigint.DivDown(bigint.Mul(product, a6), ONE_20)
	}

	if bigint.IsGreaterThanOrEqual(x, x7) {
		x = bigint.Sub(x, x7)
		product = bigint.DivDown(bigint.Mul(product, a7), ONE_20)
	}

	if bigint.IsGreaterThanOrEqual(x, x8) {
		x = bigint.Sub(x, x8)
		product = bigint.DivDown(bigint.Mul(product, a8), ONE_20)
	}

	if bigint.IsGreaterThanOrEqual(x, x9) {
		x = bigint.Sub(x, x9)
		product = bigint.DivDown(bigint.Mul(product, a9), ONE_20)
	}

	seriesSum := ONE_20
	var term *big.Int

	term = x
	seriesSum = bigint.Add(seriesSum, term)

	term = bigint.DivDown(bigint.DivDown(bigint.Mul(term, x), ONE_20), bigint.New(2))
	seriesSum = bigint.Add(seriesSum, term)

	term = bigint.DivDown(bigint.DivDown(bigint.Mul(term, x), ONE_20), bigint.New(3))
	seriesSum = bigint.Add(seriesSum, term)

	term = bigint.DivDown(bigint.DivDown(bigint.Mul(term, x), ONE_20), bigint.New(4))
	seriesSum = bigint.Add(seriesSum, term)

	term = bigint.DivDown(bigint.DivDown(bigint.Mul(term, x), ONE_20), bigint.New(5))
	seriesSum = bigint.Add(seriesSum, term)

	term = bigint.DivDown(bigint.DivDown(bigint.Mul(term, x), ONE_20), bigint.New(6))
	seriesSum = bigint.Add(seriesSum, term)

	term = bigint.DivDown(bigint.DivDown(bigint.Mul(term, x), ONE_20), bigint.New(7))
	seriesSum = bigint.Add(seriesSum, term)

	term = bigint.DivDown(bigint.DivDown(bigint.Mul(term, x), ONE_20), bigint.New(8))
	seriesSum = bigint.Add(seriesSum, term)

	term = bigint.DivDown(bigint.DivDown(bigint.Mul(term, x), ONE_20), bigint.New(9))
	seriesSum = bigint.Add(seriesSum, term)

	term = bigint.DivDown(bigint.DivDown(bigint.Mul(term, x), ONE_20), bigint.New(10))
	seriesSum = bigint.Add(seriesSum, term)

	term = bigint.DivDown(bigint.DivDown(bigint.Mul(term, x), ONE_20), bigint.New(11))
	seriesSum = bigint.Add(seriesSum, term)

	term = bigint.DivDown(bigint.DivDown(bigint.Mul(term, x), ONE_20), bigint.New(12))
	seriesSum = bigint.Add(seriesSum, term)

	return bigint.DivDown(bigint.Mul(bigint.DivDown(bigint.Mul(product, seriesSum), ONE_20), firstAN), bigint.New(100))
}

func Log(arg, base *big.Int) *big.Int {
	var logBase *big.Int
	if bigint.IsLessThan(LN_36_LOWER_BOUND, base) && bigint.IsLessThan(base, LN_36_UPPER_BOUND) {
		logBase = _ln36(base)
	} else {
		logBase = bigint.Mul(Ln(base), bigint.ONE_18)
	}

	var logArg *big.Int
	if bigint.IsLessThan(LN_36_LOWER_BOUND, arg) && bigint.IsLessThan(arg, LN_36_UPPER_BOUND) {
		logArg = _ln36(arg)
	} else {
		logArg = bigint.Mul(Ln(arg), bigint.ONE_18)
	}

	return bigint.DivDown(bigint.Mul(logArg, bigint.ONE_18), logBase)
}

func Ln(x *big.Int) *big.Int {
	if bigint.IsLessThanOrEqual(x, bigint.ZERO) {
		panic("ln error: OUT_OF_BOUNDS")
	}

	if bigint.IsLessThan(x, bigint.ONE_18) {
		return new(big.Int).Neg(Ln(bigint.DivDown(bigint.Mul(bigint.ONE_18, bigint.ONE_18), x)))
	}

	sum := bigint.ZERO
	if bigint.IsGreaterThanOrEqual(x, bigint.Mul(a0, bigint.ONE_18)) {
		x = bigint.DivDown(x, a0)
		sum = bigint.Add(sum, x0)
	}

	if bigint.IsGreaterThanOrEqual(x, bigint.Mul(a1, bigint.ONE_18)) {
		x = bigint.DivDown(x, a1)
		sum = bigint.Add(sum, x1)
	}

	sum = bigint.Mul(sum, bigint.New(100))
	x = bigint.Mul(x, bigint.New(100))

	if bigint.IsGreaterThanOrEqual(x, a2) {
		x = bigint.DivDown(bigint.Mul(x, ONE_20), a2)
		sum = bigint.Add(sum, x2)
	}

	if bigint.IsGreaterThanOrEqual(x, a3) {
		x = bigint.DivDown(bigint.Mul(x, ONE_20), a3)
		sum = bigint.Add(sum, x3)
	}

	if bigint.IsGreaterThanOrEqual(x, a4) {
		x = bigint.DivDown(bigint.Mul(x, ONE_20), a4)
		sum = bigint.Add(sum, x4)
	}

	if bigint.IsGreaterThanOrEqual(x, a5) {
		x = bigint.DivDown(bigint.Mul(x, ONE_20), a5)
		sum = bigint.Add(sum, x5)
	}

	if bigint.IsGreaterThanOrEqual(x, a6) {
		x = bigint.DivDown(bigint.Mul(x, ONE_20), a6)
		sum = bigint.Add(sum, x6)
	}

	if bigint.IsGreaterThanOrEqual(x, a7) {
		x = bigint.DivDown(bigint.Mul(x, ONE_20), a7)
		sum = bigint.Add(sum, x7)
	}

	if bigint.IsGreaterThanOrEqual(x, a8) {
		x = bigint.DivDown(bigint.Mul(x, ONE_20), a8)
		sum = bigint.Add(sum, x8)
	}

	if bigint.IsGreaterThanOrEqual(x, a9) {
		x = bigint.DivDown(bigint.Mul(x, ONE_20), a9)
		sum = bigint.Add(sum, x9)
	}

	if bigint.IsGreaterThanOrEqual(x, a10) {
		x = bigint.DivDown(bigint.Mul(x, ONE_20), a10)
		sum = bigint.Add(sum, x10)
	}

	if bigint.IsGreaterThanOrEqual(x, a11) {
		x = bigint.DivDown(bigint.Mul(x, ONE_20), a11)
		sum = bigint.Add(sum, x11)
	}

	z := bigint.DivDown(bigint.Mul(bigint.Sub(x, ONE_20), ONE_20), bigint.Add(x, ONE_20))
	z_squared := bigint.DivDown(bigint.Mul(z, z), ONE_20)

	num := z
	seriesSum := num

	num = bigint.DivDown(bigint.Mul(num, z_squared), ONE_20)
	seriesSum = bigint.Add(seriesSum, bigint.DivDown(num, bigint.New(3)))

	num = bigint.DivDown(bigint.Mul(num, z_squared), ONE_20)
	seriesSum = bigint.Add(seriesSum, bigint.DivDown(num, bigint.New(5)))

	num = bigint.DivDown(bigint.Mul(num, z_squared), ONE_20)
	seriesSum = bigint.Add(seriesSum, bigint.DivDown(num, bigint.New(7)))

	num = bigint.DivDown(bigint.Mul(num, z_squared), ONE_20)
	seriesSum = bigint.Add(seriesSum, bigint.DivDown(num, bigint.New(9)))

	num = bigint.DivDown(bigint.Mul(num, z_squared), ONE_20)
	seriesSum = bigint.Add(seriesSum, bigint.DivDown(num, bigint.New(11)))

	seriesSum = bigint.Mul(seriesSum, bigint.New(2))

	return bigint.DivDown(bigint.Add(sum, seriesSum), bigint.New(100))
}

// func _ln(x *big.Int) *big.Int {
// if bigint.IsLessThan(x, bigint.ONE_18) {
// 	return new(big.Int).Neg(_ln(bigint.DivDown(bigint.Mul(bigint.ONE_18, bigint.ONE_18), x)))
// }

// sum := bigint.ZERO
// if bigint.IsGreaterThanOrEqual(x, bigint.Mul(a0, bigint.ONE_18)) {
// 	x = bigint.DivDown(x, a0)
// 	sum = bigint.Add(sum, x0)
// }

// if bigint.IsGreaterThanOrEqual(x, bigint.Mul(a1, bigint.ONE_18)) {
// 	x = bigint.DivDown(x, a1)
// 	sum = bigint.Add(sum, x1)
// }

// sum = bigint.Mul(sum, bigint.New(100))
// x = bigint.Mul(x, bigint.New(100))

// if bigint.IsGreaterThanOrEqual(x, a2) {
// 	x = bigint.DivDown(bigint.Mul(x, ONE_20), a2)
// 	sum = bigint.Add(sum, x2)
// }

// if bigint.IsGreaterThanOrEqual(x, a3) {
// 	x = bigint.DivDown(bigint.Mul(x, ONE_20), a3)
// 	sum = bigint.Add(sum, x3)
// }

// if bigint.IsGreaterThanOrEqual(x, a4) {
// 	x = bigint.DivDown(bigint.Mul(x, ONE_20), a4)
// 	sum = bigint.Add(sum, x4)
// }

// if bigint.IsGreaterThanOrEqual(x, a5) {
// 	x = bigint.DivDown(bigint.Mul(x, ONE_20), a5)
// 	sum = bigint.Add(sum, x5)
// }

// if bigint.IsGreaterThanOrEqual(x, a6) {
// 	x = bigint.DivDown(bigint.Mul(x, ONE_20), a6)
// 	sum = bigint.Add(sum, x6)
// }

// if bigint.IsGreaterThanOrEqual(x, a7) {
// 	x = bigint.DivDown(bigint.Mul(x, ONE_20), a7)
// 	sum = bigint.Add(sum, x7)
// }

// if bigint.IsGreaterThanOrEqual(x, a8) {
// 	x = bigint.DivDown(bigint.Mul(x, ONE_20), a8)
// 	sum = bigint.Add(sum, x8)
// }

// if bigint.IsGreaterThanOrEqual(x, a9) {
// 	x = bigint.DivDown(bigint.Mul(x, ONE_20), a9)
// 	sum = bigint.Add(sum, x9)
// }

// if bigint.IsGreaterThanOrEqual(x, a10) {
// 	x = bigint.DivDown(bigint.Mul(x, ONE_20), a10)
// 	sum = bigint.Add(sum, x10)
// }

// if bigint.IsGreaterThanOrEqual(x, a11) {
// 	x = bigint.DivDown(bigint.Mul(x, ONE_20), a11)
// 	sum = bigint.Add(sum, x11)
// }

// z := bigint.DivDown(bigint.Mul(bigint.Sub(x, ONE_20), ONE_20), bigint.Add(x, ONE_20))
// z_squared := bigint.DivDown(bigint.Mul(z, z), ONE_20)

// num := z
// seriesSum := num

// num = bigint.DivDown(bigint.Mul(num, z_squared), ONE_20)
// seriesSum = bigint.Add(seriesSum, bigint.DivDown(num, bigint.New(3)))

// num = bigint.DivDown(bigint.Mul(num, z_squared), ONE_20)
// seriesSum = bigint.Add(seriesSum, bigint.DivDown(num, bigint.New(5)))

// num = bigint.DivDown(bigint.Mul(num, z_squared), ONE_20)
// seriesSum = bigint.Add(seriesSum, bigint.DivDown(num, bigint.New(7)))

// num = bigint.DivDown(bigint.Mul(num, z_squared), ONE_20)
// seriesSum = bigint.Add(seriesSum, bigint.DivDown(num, bigint.New(9)))

// num = bigint.DivDown(bigint.Mul(num, z_squared), ONE_20)
// seriesSum = bigint.Add(seriesSum, bigint.DivDown(num, bigint.New(11)))

// seriesSum = bigint.Mul(seriesSum, bigint.New(2))

// return bigint.DivDown(bigint.Add(sum, seriesSum), bigint.New(100))
// }

func _ln36(x *big.Int) *big.Int {
	x = bigint.Mul(x, bigint.ONE_18)

	z := bigint.DivDown(
		bigint.Mul(bigint.Sub(x, ONE_36), ONE_36),
		bigint.Add(x, ONE_36),
	)

	z_squared := bigint.DivDown(bigint.Mul(z, z), ONE_36)

	num := z
	seriesSum := num

	num = bigint.DivDown(bigint.Mul(num, z_squared), ONE_36)
	seriesSum = bigint.Add(seriesSum, bigint.DivDown(num, big.NewInt(3)))

	num = bigint.DivDown(bigint.Mul(num, z_squared), ONE_36)
	seriesSum = bigint.Add(seriesSum, bigint.DivDown(num, big.NewInt(5)))

	num = bigint.DivDown(bigint.Mul(num, z_squared), ONE_36)
	seriesSum = bigint.Add(seriesSum, bigint.DivDown(num, big.NewInt(7)))

	num = bigint.DivDown(bigint.Mul(num, z_squared), ONE_36)
	seriesSum = bigint.Add(seriesSum, bigint.DivDown(num, big.NewInt(9)))

	num = bigint.DivDown(bigint.Mul(num, z_squared), ONE_36)
	seriesSum = bigint.Add(seriesSum, bigint.DivDown(num, big.NewInt(11)))

	num = bigint.DivDown(bigint.Mul(num, z_squared), ONE_36)
	seriesSum = bigint.Add(seriesSum, bigint.DivDown(num, big.NewInt(13)))

	num = bigint.DivDown(bigint.Mul(num, z_squared), ONE_36)
	seriesSum = bigint.Add(seriesSum, bigint.DivDown(num, big.NewInt(15)))

	return bigint.Mul(seriesSum, big.NewInt(2))
}
