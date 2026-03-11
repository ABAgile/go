package abagile

import (
	"github.com/samber/lo"
)

// Validate is a helper that creates an error when a condition is not met.
func Validate(ok bool, format string, args ...any) error { return lo.Validate(ok, format, args...) }

// Must is a helper that wraps a call to a function returning a value and an error and panics if err is error or false.
func Must[T any](val T, err any, messageArgs ...any) T { return lo.Must(val, err, messageArgs...) }

// Must0 has the same behavior as Must, but callback returns no variable.
func Must0(err any, messageArgs ...any) { lo.Must0(err, messageArgs...) }

// Must1 is an alias to Must.
func Must1[T any](val T, err any, messageArgs ...any) T { return lo.Must1(val, err, messageArgs...) }

// Must2 has the same behavior as Must, but callback returns 2 variables.
func Must2[T1, T2 any](val1 T1, val2 T2, err any, messageArgs ...any) (T1, T2) { return lo.Must2(val1, val2, err, messageArgs...) }

// Must3 has the same behavior as Must, but callback returns 3 variables.
func Must3[T1, T2, T3 any](val1 T1, val2 T2, val3 T3, err any, messageArgs ...any) (T1, T2, T3) { return lo.Must3(val1, val2, val3, err, messageArgs...) }

// Must4 has the same behavior as Must, but callback returns 4 variables.
func Must4[T1, T2, T3, T4 any](val1 T1, val2 T2, val3 T3, val4 T4, err any, messageArgs ...any) (T1, T2, T3, T4) { return lo.Must4(val1, val2, val3, val4, err, messageArgs...) }

// Must5 has the same behavior as Must, but callback returns 5 variables.
func Must5[T1, T2, T3, T4, T5 any](val1 T1, val2 T2, val3 T3, val4 T4, val5 T5, err any, messageArgs ...any) (T1, T2, T3, T4, T5) { return lo.Must5(val1, val2, val3, val4, val5, err, messageArgs...) }

// Must6 has the same behavior as Must, but callback returns 6 variables.
func Must6[T1, T2, T3, T4, T5, T6 any](val1 T1, val2 T2, val3 T3, val4 T4, val5 T5, val6 T6, err any, messageArgs ...any) (T1, T2, T3, T4, T5, T6) { return lo.Must6(val1, val2, val3, val4, val5, val6, err, messageArgs...) }

// Try calls the function and return false in case of error.
func Try(callback func() error) bool { return lo.Try(callback) }

// Try0 has the same behavior as Try, but callback returns no variable.
func Try0(callback func()) bool { return lo.Try0(callback) }

// Try1 is an alias to Try.
func Try1(callback func() error) bool { return lo.Try1(callback) }

// Try2 has the same behavior as Try, but callback returns 2 variables.
func Try2[T any](callback func() (T, error)) bool { return lo.Try2(callback) }

// Try3 has the same behavior as Try, but callback returns 3 variables.
func Try3[T, R any](callback func() (T, R, error)) bool { return lo.Try3(callback) }

// Try4 has the same behavior as Try, but callback returns 4 variables.
func Try4[T, R, S any](callback func() (T, R, S, error)) bool { return lo.Try4(callback) }

// Try5 has the same behavior as Try, but callback returns 5 variables.
func Try5[T, R, S, Q any](callback func() (T, R, S, Q, error)) bool { return lo.Try5(callback) }

// Try6 has the same behavior as Try, but callback returns 6 variables.
func Try6[T, R, S, Q, U any](callback func() (T, R, S, Q, U, error)) bool { return lo.Try6(callback) }

// TryOr has the same behavior as Must, but returns a default value in case of error.
func TryOr[A any](callback func() (A, error), fallbackA A) (A, bool) { return lo.TryOr(callback, fallbackA) }

// TryOr1 has the same behavior as Must, but returns a default value in case of error.
func TryOr1[A any](callback func() (A, error), fallbackA A) (A, bool) { return lo.TryOr1(callback, fallbackA) }

// TryOr2 has the same behavior as Must, but returns a default value in case of error.
func TryOr2[A, B any](callback func() (A, B, error), fallbackA A, fallbackB B) (A, B, bool) { return lo.TryOr2(callback, fallbackA, fallbackB) }

// TryOr3 has the same behavior as Must, but returns a default value in case of error.
func TryOr3[A, B, C any](callback func() (A, B, C, error), fallbackA A, fallbackB B, fallbackC C) (A, B, C, bool) { return lo.TryOr3(callback, fallbackA, fallbackB, fallbackC) }

// TryOr4 has the same behavior as Must, but returns a default value in case of error.
func TryOr4[A, B, C, D any](callback func() (A, B, C, D, error), fallbackA A, fallbackB B, fallbackC C, fallbackD D) (A, B, C, D, bool) { return lo.TryOr4(callback, fallbackA, fallbackB, fallbackC, fallbackD) }

// TryOr5 has the same behavior as Must, but returns a default value in case of error.
func TryOr5[A, B, C, D, E any](callback func() (A, B, C, D, E, error), fallbackA A, fallbackB B, fallbackC C, fallbackD D, fallbackE E) (A, B, C, D, E, bool) { return lo.TryOr5(callback, fallbackA, fallbackB, fallbackC, fallbackD, fallbackE) }

// TryOr6 has the same behavior as Must, but returns a default value in case of error.
func TryOr6[A, B, C, D, E, F any](callback func() (A, B, C, D, E, F, error), fallbackA A, fallbackB B, fallbackC C, fallbackD D, fallbackE E, fallbackF F) (A, B, C, D, E, F, bool) { return lo.TryOr6(callback, fallbackA, fallbackB, fallbackC, fallbackD, fallbackE, fallbackF) }

// TryWithErrorValue has the same behavior as Try, but also returns value passed to panic.
func TryWithErrorValue(callback func() error) (any, bool) { return lo.TryWithErrorValue(callback) }

// TryCatch has the same behavior as Try, but calls the catch function in case of error.
func TryCatch(callback func() error, catch func()) { lo.TryCatch(callback, catch) }

// TryCatchWithErrorValue has the same behavior as TryWithErrorValue, but calls the catch function in case of error.
func TryCatchWithErrorValue(callback func() error, catch func(any)) { lo.TryCatchWithErrorValue(callback, catch) }

// ErrorsAs is a shortcut for errors.As(err, &&T).
func ErrorsAs[T error](err error) (T, bool) { return lo.ErrorsAs[T](err) }

// Assert does nothing when the condition is true, otherwise it panics with an optional message.
func Assert(condition bool, message ...string) { lo.Assert(condition, message...) }

// Assertf does nothing when the condition is true, otherwise it panics with a formatted message.
func Assertf(condition bool, format string, args ...any) { lo.Assertf(condition, format, args...) }
