package stacks

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		input string
		exp   string
	}{
		{"g", "g"},
		{"ggaabcdeb", "bcdeb"},
		{"abbddaccaaabcd", "abcd"},
		{"aabbccdd", ""},
		{"aannkwwwkkkwna", "kwkwna"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			if result := removeDuplicates(tc.input); result != tc.exp {
				t.Errorf("exp: %v, got: %v", tc.exp, result)
			}
		})
	}
}

func BenchmarkRemoveDuplicates(b *testing.B) {
	str := "aannkwwddowwlaammdkwkkkwna"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = removeDuplicates(str)
	}
}

/*func BenchmarkRemoveDuplicatesWithPrev(b *testing.B) {
	str := "aannkwwddowwlaammdkwkkkwna"
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = removeDuplicatesWithPrev(str)
	}
}
*/
