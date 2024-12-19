//go:build go1.23

package sets

import (
	"iter"
)

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
