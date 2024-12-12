//go:build go1.23

package sets

import (
	"iter"
)

// Values iterator over sequences of individual values.
//
// Deprecated: use All instead.
func (s Set[T]) Values() iter.Seq[T] { return s.All() }

// All iterator over sequences of individual values.
func (s Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for k := range s {
			if !yield(k) {
				return
			}
		}
	}
}
