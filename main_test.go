package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		result int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{3, 4, 7},
	}

	for _, test := range tests {
		res := add(test.a, test.b)
		if res != test.result {
			t.Errorf("add(%d, %d) = %d, want %d", test.a, test.b, res, test.result)
		}
	}
}

func TestSub(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		result int
	}{
		{1, 2, -1},
		{10, 3, 7},
		{3, 1, 2},
	}

	for _, test := range tests {
		res := sub(test.a, test.b)
		if res != test.result {
			t.Errorf("sub(%d, %d) = %d, want %d", test.a, test.b, res, test.result)
		}
	}
}

func TestMul(t *testing.T) {
	tests := []struct {
		a      int
		b      int
		result int
	}{
		{1, 2, 2},
		{10, 3, 30},
		{3, 5, 15},
	}

	for _, test := range tests {
		res := mul(test.a, test.b)
		if res != test.result {
			t.Errorf("mul(%d, %d) = %d, want %d", test.a, test.b, res, test.result)
		}
	}
}
