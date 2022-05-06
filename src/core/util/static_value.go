package util

import "math"

const LONG_MAX_VALUE = math.MaxInt64

type All interface {
	any | int64 | float64 | int | string | int32
}

type MyInt interface {
	any | int
}
