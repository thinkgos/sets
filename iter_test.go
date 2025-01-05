//go:build go1.23

package sets

import "testing"

func Test_Iter_Values(t *testing.T) {
	s := New(1, 3, 4, 6)
	for _, v := range s.Values() {
		t.Log(v)
	}
}

func Test_Iter_Values_Return(t *testing.T) {
	s := New(1, 3, 4, 6)
	for _, v := range s.Values() {
		t.Log(v)
		if v == 3 {
			return
		}
	}
}
