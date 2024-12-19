package sets

// Set sets.Set is a set of `T`,
// implemented via map[T]struct{} for minimal memory consumption.
type Set[T comparable] map[T]struct{}

func newSet[T comparable](cap int) Set[T] {
	return make(Set[T], cap)
}

// New creates a T from a list of values.
func New[T comparable](items ...T) Set[T] {
	ret := newSet[T](len(items))
	return ret.Insert(items...)
}

// NewFrom creates a T from a keys of a map[T](? extends any).
// If the value passed in is not actually a map, this will panic.
func NewFrom[T comparable, V any, M ~map[T]V](m M) Set[T] {
	ret := newSet[T](len(m))
	for k := range m {
		ret[k] = struct{}{}
	}
	return ret
}

// Insert adds items to the set.
func (s Set[T]) Insert(items ...T) Set[T] {
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

// Delete removes all items from the set.
func (s Set[T]) Delete(items ...T) Set[T] {
	for _, item := range items {
		delete(s, item)
	}
	return s
}

// Contains returns true if and only if item is contained in the set.
func (s Set[T]) Contains(item T) bool {
	_, contained := s[item]
	return contained
}

// ContainsAll returns true if and only if all items are contained in the set.
func (s Set[T]) ContainsAll(items ...T) bool {
	for _, item := range items {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}

// ContainsAny returns true if any items are contained in the set.
func (s Set[T]) ContainsAny(items ...T) bool {
	for _, item := range items {
		if s.Contains(item) {
			return true
		}
	}
	return false
}

// Merge is like Union, however it modifies the current set it's applied on
// with the given s2 set.
// For example:
// s1 = {a1, a2}
// s2 = {a3, a4}
// s1.Merge(s2), s1 = {a1, a2, a3, a4}
// s2.Merge(s1), s2 = {a1, a2, a3, a4}.
func (s Set[T]) Merge(s2 Set[T]) Set[T] {
	for item := range s2 {
		s[item] = struct{}{}
	}
	return s
}

// IsSuperset returns true if and only if s1 is a superset of s2.
func (s Set[T]) IsSuperset(s2 Set[T]) bool {
	for item := range s2 {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}

// IsSubset returns true if and only if s1 is a superset of s2.
func (s Set[T]) IsSubset(s2 Set[T]) bool {
	for item := range s {
		if !s2.Contains(item) {
			return false
		}
	}
	return true
}

// Equal returns true if and only if s1 is equal (as a set) to s2.
// Two sets are equal if their membership is identical.
// (In practice, this means same elements, order doesn't matter).
func (s Set[T]) Equal(s2 Set[T]) bool {
	return len(s) == len(s2) && s.IsSuperset(s2)
}

// Pop Returns a single element from the set.
func (s Set[T]) Pop() (v T, ok bool) {
	for key := range s {
		delete(s, key)
		return key, true
	}
	return
}

// Len returns the size of the set.
func (s Set[T]) Len() int { return len(s) }

// Clone returns a new Set with a copy of s.
func (s Set[T]) Clone() Set[T] {
	return NewFrom(s)
}

// Values returns the value of the sets.
func (s Set[T]) Values() []T {
	res := make([]T, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	return res
}

// List returns the value of the sets.
//
// Deprecated: use Values instead.
func (s Set[T]) List() []T {
	return s.Values()
}

// Each traverses the items in the Set, calling the provided function for each
// set member. Traversal will continue until all items in the Set have been
// visited, or if the closure returns false.
func (s Set[T]) Each(f func(item T) bool) {
	for item := range s {
		if !f(item) {
			break
		}
	}
}
