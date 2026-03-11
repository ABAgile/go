package abagile

import "github.com/samber/lo"

// IsNil checks if a value is nil or if it's a reference type with a nil underlying value.
func IsNil(x any) bool { return lo.IsNil(x) }

// IsNotNil checks if a value is not nil or if it's not a reference type with a nil underlying value.
func IsNotNil(x any) bool { return !lo.IsNil(x) }

// Nil returns a nil pointer of type.
func Nil[T any]() *T { return lo.Nil[T]() }

// EmptyableToPtr returns a pointer copy of value if it's nonzero. Otherwise, returns nil pointer.
func EmptyableToPtr[T any](x T) *T { return lo.EmptyableToPtr(x) }

// FromPtr returns the pointer value or empty.
func FromPtr[T any](x *T) T { return lo.FromPtr(x) }

// FromPtrOr returns the pointer value or the fallback value.
func FromPtrOr[T any](x *T, fallback T) T { return lo.FromPtrOr(x, fallback) }

// ToSlicePtr returns a slice of pointers to each value.
func ToSlicePtr[T any](collection []T) []*T { return lo.ToSlicePtr(collection) }

// FromSlicePtr returns a slice with the pointer values. Returns a zero value in case of a nil pointer element.
func FromSlicePtr[T any](collection []*T) []T { return lo.FromSlicePtr(collection) }

// FromSlicePtrOr returns a slice with the pointer values or the fallback value.
func FromSlicePtrOr[T any](collection []*T, fallback T) []T {
	return lo.FromSlicePtrOr(collection, fallback)
}

// ToAnySlice returns a slice with all elements mapped to `any` type.
func ToAnySlice[T any](collection []T) []any { return lo.ToAnySlice(collection) }

// FromAnySlice returns a slice with all elements mapped to a type. Returns false in case of type conversion failure.
func FromAnySlice[T any](in []any) ([]T, bool) { return lo.FromAnySlice[T](in) }

// Empty returns the zero value.
func Empty[T any]() T { return lo.Empty[T]() }

// IsEmpty returns true if argument is a zero value.
func IsEmpty[T comparable](v T) bool { return lo.IsEmpty(v) }

// IsNotEmpty returns true if argument is not a zero value.
func IsNotEmpty[T comparable](v T) bool { return lo.IsNotEmpty(v) }

// Coalesce returns the first non-empty arguments. Arguments must be comparable.
func Coalesce[T comparable](values ...T) (T, bool) { return lo.Coalesce(values...) }

// CoalesceOrEmpty returns the first non-empty arguments. Arguments must be comparable.
func CoalesceOrEmpty[T comparable](v ...T) T { return lo.CoalesceOrEmpty(v...) }

// CoalesceSlice returns the first non-zero slice.
func CoalesceSlice[T any](v ...[]T) ([]T, bool) { return lo.CoalesceSlice(v...) }

// CoalesceSliceOrEmpty returns the first non-zero slice.
func CoalesceSliceOrEmpty[T any](v ...[]T) []T { return lo.CoalesceSliceOrEmpty(v...) }

// CoalesceMap returns the first non-zero map.
func CoalesceMap[K comparable, V any](v ...map[K]V) (map[K]V, bool) { return lo.CoalesceMap(v...) }

// CoalesceMapOrEmpty returns the first non-zero map.
func CoalesceMapOrEmpty[K comparable, V any](v ...map[K]V) map[K]V {
	return lo.CoalesceMapOrEmpty(v...)
}
