package abagile

import "github.com/samber/lo"

var (
	LowerCaseLettersCharset = lo.LowerCaseLettersCharset
	UpperCaseLettersCharset = lo.UpperCaseLettersCharset
	LettersCharset          = lo.LettersCharset
	NumbersCharset          = lo.NumbersCharset
	AlphanumericCharset     = lo.AlphanumericCharset
	SpecialCharset          = lo.SpecialCharset
	AllCharset              = lo.AllCharset
)

// RandomString return a random string.
func RandomString(size int, charset []rune) string { return lo.RandomString(size, charset) }

// Substring extracts a substring from a string with Unicode character (rune) awareness.
func Substring[T ~string](str T, offset int, length uint) T { return lo.Substring(str, offset, length) }

// ChunkString returns a slice of strings split into groups of length size.
func ChunkString[T ~string](str T, size int) []T { return lo.ChunkString(str, size) }

// RuneLength is an alias to utf8.RuneCountInString which returns the number of runes in string.
func RuneLength(str string) int { return lo.RuneLength(str) }

// PascalCase converts string to pascal case.
func PascalCase(str string) string { return lo.PascalCase(str) }

// CamelCase converts string to camel case.
func CamelCase(str string) string { return lo.CamelCase(str) }

// KebabCase converts string to kebab case.
func KebabCase(str string) string { return lo.KebabCase(str) }

// SnakeCase converts string to snake case.
func SnakeCase(str string) string { return lo.SnakeCase(str) }

// Words splits string into a slice of its words.
func Words(str string) []string { return lo.Words(str) }

// Capitalize converts the first character of string to upper case and the remaining to lower case.
func Capitalize(str string) string { return lo.Capitalize(str) }

// Ellipsis trims and truncates a string to a specified length in runes and appends an ellipsis if truncated.
func Ellipsis(str string, length int) string { return lo.Ellipsis(str, length) }
