package fruitsIntoBaskets

import (
	"fmt"
	"testing"
)

func TestMaxFruitCountOf2Types(t *testing.T) {
	testCases := []struct {
		input    []byte
		expected int
	}{
		{[]byte{'A', 'B', 'C', 'A', 'C'}, 3},
		{[]byte{'A', 'B', 'C', 'B', 'B', 'C'}, 5},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("string: %v", tc.input), func(t *testing.T) {
			result := maxFruitCountOf2Types(tc.input)

			if tc.expected != result {
				t.Fatalf("Exp result %d, got %d", tc.expected, result)
			}
		})
	}
}
