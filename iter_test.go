//go:build go1.23

package sets

import "testing"

func Test_Iter_Values(t *testing.T) {
	s := New(1, 3, 4, 6)
	for k := range s.Values() {
		t.Log(k)
	}
}

func Test_Iter_Values_Return(t *testing.T) {
	s := New(1, 3, 4, 6)
	for k := range s.Values() {
		t.Log(k)
		if k == 3 {
			return
		}
	}
}
