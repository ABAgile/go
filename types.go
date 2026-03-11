package abagile

import "github.com/samber/lo"

// Entry defines a key/value pairs.
type Entry[K comparable, V any] lo.Entry[K, V]

// Tuple2 is a group of 2 elements (pair).
type Tuple2[A, B any] lo.Tuple2[A, B]

// Unpack returns values contained in a tuple.
func (t Tuple2[A, B]) Unpack() (A, B) { return t.A, t.B }

// Tuple3 is a group of 3 elements.
type Tuple3[A, B, C any] lo.Tuple3[A, B, C]

// Unpack returns values contained in a tuple.
func (t Tuple3[A, B, C]) Unpack() (A, B, C) { return t.A, t.B, t.C }

// Tuple4 is a group of 4 elements.
type Tuple4[A, B, C, D any] lo.Tuple4[A, B, C, D]

// Unpack returns values contained in a tuple.
func (t Tuple4[A, B, C, D]) Unpack() (A, B, C, D) { return t.A, t.B, t.C, t.D }

// Tuple5 is a group of 5 elements.
type Tuple5[A, B, C, D, E any] lo.Tuple5[A, B, C, D, E]

// Unpack returns values contained in a tuple.
func (t Tuple5[A, B, C, D, E]) Unpack() (A, B, C, D, E) { return t.A, t.B, t.C, t.D, t.E }

// Tuple6 is a group of 6 elements.
type Tuple6[A, B, C, D, E, F any] lo.Tuple6[A, B, C, D, E, F]

// Unpack returns values contained in a tuple.
func (t Tuple6[A, B, C, D, E, F]) Unpack() (A, B, C, D, E, F) { return t.A, t.B, t.C, t.D, t.E, t.F }

// Tuple7 is a group of 7 elements.
type Tuple7[A, B, C, D, E, F, G any] lo.Tuple7[A, B, C, D, E, F, G]

// Unpack returns values contained in a tuple.
func (t Tuple7[A, B, C, D, E, F, G]) Unpack() (A, B, C, D, E, F, G) { return t.A, t.B, t.C, t.D, t.E, t.F, t.G }

// Tuple8 is a group of 8 elements.
type Tuple8[A, B, C, D, E, F, G, H any] lo.Tuple8[A, B, C, D, E, F, G, H]

// Unpack returns values contained in a tuple.
func (t Tuple8[A, B, C, D, E, F, G, H]) Unpack() (A, B, C, D, E, F, G, H) { return t.A, t.B, t.C, t.D, t.E, t.F, t.G, t.H }

// Tuple9 is a group of 9 elements.
type Tuple9[A, B, C, D, E, F, G, H, I any] lo.Tuple9[A, B, C, D, E, F, G, H, I]

// Unpack returns values contained in a tuple.
func (t Tuple9[A, B, C, D, E, F, G, H, I]) Unpack() (A, B, C, D, E, F, G, H, I) { return t.A, t.B, t.C, t.D, t.E, t.F, t.G, t.H, t.I }
