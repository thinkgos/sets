package sets

import (
	"reflect"
	"slices"
	"sort"
	"testing"
)

func TestSetDifference(t *testing.T) {
	a := New(1, 2, 3)
	b := New(1, 2, 4, 5)

	t.Run("set", func(t *testing.T) {
		c := a.Difference(b)
		d := b.Difference(a)
		if len(c) != 1 {
			t.Errorf("Expected len=1: %d", len(c))
		}
		if !c.Contains(3) {
			t.Errorf("Unexpected contents: %#v", c.List())
		}
		if len(d) != 2 {
			t.Errorf("Expected len=2: %d", len(d))
		}
		if !d.Contains(4) || !d.Contains(5) {
			t.Errorf("Unexpected contents: %#v", d.List())
		}
	})

	t.Run("slice", func(t *testing.T) {
		c := a.DifferenceSlice(b)
		d := b.DifferenceSlice(a)
		if len(c) != 1 {
			t.Errorf("Expected len=1: %d", len(c))
		}
		if !slices.Contains(c, 3) {
			t.Errorf("Unexpected contents: %#v", c)
		}
		if len(d) != 2 {
			t.Errorf("Expected len=2: %d", len(d))
		}
		if !slices.Contains(d, 4) || !slices.Contains(d, 5) {
			t.Errorf("Unexpected contents: %#v", d)
		}
	})
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		s1       Set[int]
		s2       Set[int]
		expected Set[int]
	}{
		{
			New[int](1, 2, 3, 4),
			New[int](3, 4, 5, 6),
			New[int](3, 4),
		},
		{
			New[int](1, 2, 3, 4),
			New[int](1, 2, 3, 4),
			New[int](1, 2, 3, 4),
		},
		{
			New[int](1, 2, 3, 4),
			New[int](),
			New[int](),
		},
		{
			New[int](),
			New[int](1, 2, 3, 4),
			New[int](),
		},
		{
			New[int](),
			New[int](),
			New[int](),
		},
	}

	for _, test := range tests {
		// sets
		intersection := test.s1.Intersection(test.s2)
		if intersection.Len() != test.expected.Len() {
			t.Errorf("Expected intersection.Len()=%d but got %d", test.expected.Len(), intersection.Len())
		}
		if !intersection.Equal(test.expected) {
			t.Errorf("Expected intersection.Equal(expected) but not true.  intersection:%v expected:%v",
				intersection.List(), test.expected.List())
		}

		// slice
		intersectionSlice := test.s1.IntersectionSlice(test.s2)
		if len(intersectionSlice) != test.expected.Len() {
			t.Errorf("Expected IntersectionSlice length =%d but got %d", test.expected.Len(), len(intersectionSlice))
		}
		if !New(intersectionSlice...).Equal(test.expected) {
			t.Errorf("Expected intersection.Equal(expected) but not true.  intersection:%v expected:%v",
				intersectionSlice, test.expected.List())
		}
	}
}
func TestUnion(t *testing.T) {
	tests := []struct {
		s1       Set[int]
		s2       Set[int]
		expected Set[int]
	}{
		{
			New[int](1, 2, 3, 4),
			New[int](3, 4, 5, 6),
			New[int](1, 2, 3, 4, 5, 6),
		},
		{
			New[int](1, 2, 3, 4),
			New[int](),
			New[int](1, 2, 3, 4),
		},
		{
			New[int](),
			New[int](1, 2, 3, 4),
			New[int](1, 2, 3, 4),
		},
		{
			New[int](),
			New[int](),
			New[int](),
		},
	}

	for _, test := range tests {
		// set
		union := test.s1.Union(test.s2)
		if union.Len() != test.expected.Len() {
			t.Errorf("Expected union.Len()=%d but got %d", test.expected.Len(), union.Len())
		}
		if !union.Equal(test.expected) {
			t.Errorf("Expected union.Equal(expected) but not true.  union:%v expected:%v", union.List(), test.expected.List())
		}
		// slice
		unionSlices := test.s1.UnionSlice(test.s2)
		if len(unionSlices) != test.expected.Len() {
			t.Errorf("Expected UnionSlices length=%d but got %d", test.expected.Len(), len(unionSlices))
		}
		if !New(unionSlices...).Equal(test.expected) {
			t.Errorf("Expected union.Equal(expected) but not true.  union:%v expected:%v", unionSlices, test.expected.List())
		}
	}
}

func TestSetDiff(t *testing.T) {
	t.Run("set", func(t *testing.T) {
		a := New(1, 3, 5, 7)
		b := New(3, 4, 5, 6)

		added, removed, remained := a.Diff(b)
		added1, removed1 := a.DiffVary(b)

		wantAdded := Set[int]{
			4: {},
			6: {},
		}
		if !reflect.DeepEqual(added, wantAdded) {
			t.Errorf("Expected %#v, Got: %#v", wantAdded, added)
		}
		if !reflect.DeepEqual(added1, wantAdded) {
			t.Errorf("Expected %#v, Got: %#v", wantAdded, added1)
		}

		wantRemoved := Set[int]{
			1: {},
			7: {},
		}
		if !reflect.DeepEqual(removed, wantRemoved) {
			t.Errorf("Expected %v, Got: %v", wantRemoved, removed)
		}
		if !reflect.DeepEqual(removed1, wantRemoved) {
			t.Errorf("Expected %v, Got: %v", wantRemoved, removed1)
		}

		wantRemained := Set[int]{
			3: {},
			5: {},
		}
		if !reflect.DeepEqual(remained, wantRemained) {
			t.Errorf("Expected %v, Got: %v", wantRemained, remained)
		}
	})
	t.Run("slice", func(t *testing.T) {
		a := New(1, 3, 5, 7)
		b := New(3, 4, 5, 6)

		added, removed, remained := a.DiffSlice(b)
		added1, removed1 := a.DiffVarySlice(b)

		sort.Ints(added)
		sort.Ints(removed)
		sort.Ints(remained)
		sort.Ints(added1)
		sort.Ints(removed1)

		wantAdded := []int{4, 6}
		if !reflect.DeepEqual(added, wantAdded) {
			t.Errorf("Expected %#v, Got: %#v", wantAdded, added)
		}
		if !reflect.DeepEqual(added1, wantAdded) {
			t.Errorf("Expected %#v, Got: %#v", wantAdded, added1)
		}

		wantRemoved := []int{1, 7}
		if !reflect.DeepEqual(removed, wantRemoved) {
			t.Errorf("Expected %v, Got: %v", wantRemoved, removed)
		}
		if !reflect.DeepEqual(removed1, wantRemoved) {
			t.Errorf("Expected %v, Got: %v", wantRemoved, removed1)
		}

		wantRemained := []int{3, 5}
		if !reflect.DeepEqual(remained, wantRemained) {
			t.Errorf("Expected %v, Got: %v", wantRemained, remained)
		}
	})
}
