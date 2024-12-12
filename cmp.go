package sets

// Difference returns a set of objects that are not in s2
// For example:
// s1 = {a1, a2, a3}
// s2 = {a1, a2, a4, a5}
// s1.Difference(s2) = {a3}
// s2.Difference(s1) = {a4, a5}.
func (s Set[T]) Difference(s2 Set[T]) Set[T] {
	result := newSet[T](len(s))
	for key := range s {
		if !s2.Contains(key) {
			result[key] = struct{}{}
		}
	}
	return result
}

// Union returns a new set which includes items in either s1 or s2.
// For example:
// s1 = {a1, a2}
// s2 = {a3, a4}
// s1.Union(s2) = {a1, a2, a3, a4}
// s2.Union(s1) = {a1, a2, a3, a4}.
func (s Set[T]) Union(s2 Set[T]) Set[T] {
	result := newSet[T](len(s) + len(s2))
	for key := range s {
		result[key] = struct{}{}
	}
	for key := range s2 {
		result[key] = struct{}{}
	}
	return result
}

// Intersection returns a new set which includes the item in BOTH s1 and s2
// For example:
// s1 = {a1, a2}
// s2 = {a2, a3}
// s1.Intersection(s2) = {a2}.
func (s Set[T]) Intersection(s2 Set[T]) Set[T] {
	var walk, other Set[T]

	if s.Len() < s2.Len() {
		walk = s
		other = s2
	} else {
		walk = s2
		other = s
	}
	ret := newSet[T](min(len(s), len(s2)))
	for key := range walk {
		if other.Contains(key) {
			ret[key] = struct{}{}
		}
	}
	return ret
}

// Diff returns s diff of s2, return added, removed, remained sets
// with the given s2 set.
// For example:
// s1 = {a1, a3, a5, a7}
// s2 = {a3, a4, a5, a6}
// added = {a4, a6}
// removed = {a1, a7}
// remained = {a3, a6}
func (s Set[T]) Diff(s2 Set[T]) (added, removed, remained Set[T]) {
	removed = newSet[T](len(s))
	added = newSet[T](len(s2))
	remained = newSet[T](len(s))
	for key := range s {
		if s2.Contains(key) {
			remained[key] = struct{}{}
		} else {
			removed[key] = struct{}{}
		}
	}
	for key := range s2 {
		if !s.Contains(key) {
			added[key] = struct{}{}
		}
	}
	return added, removed, remained
}

// DiffVary returns s diff of s2, return added, removed sets
// with the given s2 set.
// For example:
// s1 = {a1, a3, a5, a7}
// s2 = {a3, a4, a5, a6}
// added = {a4, a6}
// removed = {a1, a7}
func (s Set[T]) DiffVary(s2 Set[T]) (added, removed Set[T]) {
	removed = newSet[T](len(s))
	added = newSet[T](len(s2))
	for key := range s {
		if !s2.Contains(key) {
			removed[key] = struct{}{}
		}
	}
	for key := range s2 {
		if !s.Contains(key) {
			added[key] = struct{}{}
		}
	}
	return added, removed
}

//* diff slices

// DifferenceSlice returns a slices of objects that are not in s2
// For example:
// s1 = {a1, a2, a3}
// s2 = {a1, a2, a4, a5}
// s1.Difference(s2) = {a3}
// s2.Difference(s1) = {a4, a5}.
func (s Set[T]) DifferenceSlice(s2 Set[T]) []T {
	result := make([]T, 0, len(s))
	for key := range s {
		if !s2.Contains(key) {
			result = append(result, key)
		}
	}
	return result
}

// UnionSlice returns a slice which includes items in either s1 or s2.
// For example:
// s1 = {a1, a2}
// s2 = {a3, a4}
// s1.Union(s2) = {a1, a2, a3, a4}
// s2.Union(s1) = {a1, a2, a3, a4}.
func (s Set[T]) UnionSlice(s2 Set[T]) []T {
	result := make([]T, 0, len(s)+len(s2))
	for key := range s {
		result = append(result, key)
	}
	for key := range s2 {
		if !s.Contains(key) {
			result = append(result, key)
		}
	}
	return result
}

// IntersectionSlice returns a slice which includes the item in BOTH s1 and s2
// For example:
// s1 = {a1, a2}
// s2 = {a2, a3}
// s1.Intersection(s2) = {a2}.
func (s Set[T]) IntersectionSlice(s2 Set[T]) []T {
	var walk, other Set[T]

	if s.Len() < s2.Len() {
		walk = s
		other = s2
	} else {
		walk = s2
		other = s
	}
	ret := make([]T, 0, min(len(s), len(s2)))
	for key := range walk {
		if other.Contains(key) {
			ret = append(ret, key)
		}
	}
	return ret
}

// DiffSlice returns s diff of s2, return added, removed, remained slices
// with the given s2 set.
// For example:
// s1 = {a1, a3, a5, a7}
// s2 = {a3, a4, a5, a6}
// added = {a4, a6}
// removed = {a1, a7}
// remained = {a3, a6}
func (s Set[T]) DiffSlice(s2 Set[T]) (added, removed, remained []T) {
	removed = make([]T, 0, len(s))
	added = make([]T, 0, len(s2))
	remained = make([]T, 0, len(s))
	for key := range s {
		if s2.Contains(key) {
			remained = append(remained, key)
		} else {
			removed = append(removed, key)
		}
	}
	for key := range s2 {
		if !s.Contains(key) {
			added = append(added, key)
		}
	}
	return added, removed, remained
}

// DiffVarySlice returns s diff of s2, return added, removed slices
// with the given s2 set.
// For example:
// s1 = {a1, a3, a5, a7}
// s2 = {a3, a4, a5, a6}
// added = {a4, a6}
// removed = {a1, a7}
func (s Set[T]) DiffVarySlice(s2 Set[T]) (added, removed []T) {
	removed = make([]T, 0, len(s))
	added = make([]T, 0, len(s2))
	for key := range s {
		if !s2.Contains(key) {
			removed = append(removed, key)
		}
	}
	for key := range s2 {
		if !s.Contains(key) {
			added = append(added, key)
		}
	}
	return added, removed
}
