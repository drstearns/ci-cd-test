package main

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		a        int
		b        int
		expected int
	}{
		{0, 0, 0},
		{1, 2, 3},
	}
	for _, c := range cases {
		actual := add(c.a, c.b)
		if actual != c.expected {
			t.Errorf("expected %d but got %d", c.expected, actual)
		}
	}
}
