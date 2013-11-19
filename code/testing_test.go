// +build OMIT
package testtest

import "testing"

// START0 OMIT
func Sum(x, y int) int {
	return x + y
}

// STOP0 OMIT

// START1 OMIT
func TestSum(t *testing.T) {
	testMap := [][]int{
		{1, 1, 2},
		{2, 2, 4},
		{4, 4, 8},
		{5, 15, 20},
	}

	for _, v := range testMap {
		if i := Sum(v[0], v[1]); i != v[2] {
			t.Errorf("Error at Sum(%d, %d) returned %d", v[0], v[1], i)
		}
	}
}

// STOP1 OMIT
