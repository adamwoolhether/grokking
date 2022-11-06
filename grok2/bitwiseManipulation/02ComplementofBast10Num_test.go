package bitwiseManipulation

import (
	"fmt"
	"testing"
)

func TestFindBitwiseComplement(t *testing.T) {
	testCases := []struct {
		num int
		exp int
	}{
		{42, 21},
		{8, 7},
		{233, 22},
		{100, 27},
		{39, 24},
		{11111, 5272},
		{34631, 30904},
		{999999, 48576},
		{54, 9},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d", tc.num), func(t *testing.T) {
			if result := findBitwiseComplement(tc.num); result != tc.exp {
				t.Errorf("exp: %d, got: %d", tc.exp, result)
			}
		})
	}
}
