package abagile

import (
	"github.com/samber/lo"
)

// Range creates a slice of numbers (positive and/or negative) with given length.
func Range(elementNum int) []int { return lo.Range(elementNum) }

// RangeFrom creates a slice of numbers from start with specified length.
func RangeFrom[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](start T, elementNum int) []T { return lo.RangeFrom(start, elementNum) }

// RangeWithSteps creates a slice of numbers (positive and/or negative) progressing from start up to, but not including end.
func RangeWithSteps[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](start, end, step T) []T { return lo.RangeWithSteps(start, end, step) }

// Clamp clamps number within the inclusive lower and upper bounds.
func Clamp[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string](value, mIn, mAx T) T { return lo.Clamp(value, mIn, mAx) }

// Sum sums the values in a collection. If collection is empty 0 is returned.
func Sum[T float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | complex64 | complex128](collection []T) T { return lo.Sum(collection) }

// SumBy summarizes the values in a collection using the given return value from the iteration function.
func SumBy[T any, R float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | complex64 | complex128](collection []T, iteratee func(item T) R) R { return lo.SumBy(collection, iteratee) }

// SumByErr summarizes the values in a collection using the given return value from the iteration function.
func SumByErr[T any, R float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | complex64 | complex128](collection []T, iteratee func(item T) (R, error)) (R, error) { return lo.SumByErr(collection, iteratee) }

// Product gets the product of the values in a collection. If collection is empty 1 is returned.
func Product[T float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | complex64 | complex128](collection []T) T { return lo.Product(collection) }

// ProductBy summarizes the values in a collection using the given return value from the iteration function.
func ProductBy[T any, R float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | complex64 | complex128](collection []T, iteratee func(item T) R) R { return lo.ProductBy(collection, iteratee) }

// ProductByErr summarizes the values in a collection using the given return value from the iteration function.
func ProductByErr[T any, R float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | complex64 | complex128](collection []T, iteratee func(item T) (R, error)) (R, error) { return lo.ProductByErr(collection, iteratee) }

// Mean calculates the mean of a collection of numbers.
func Mean[T float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64](collection []T) T { return lo.Mean(collection) }

// MeanBy calculates the mean of a collection of numbers using the given return value from the iteration function.
func MeanBy[T any, R float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64](collection []T, iteratee func(item T) R) R { return lo.MeanBy(collection, iteratee) }

// MeanByErr calculates the mean of a collection of numbers using the given return value from the iteration function.
func MeanByErr[T any, R float32 | float64 | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64](collection []T, iteratee func(item T) (R, error)) (R, error) { return lo.MeanByErr(collection, iteratee) }

// Mode returns the mode (most frequent value) of a collection.
func Mode[T int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64](collection []T) []T { return lo.Mode(collection) }
