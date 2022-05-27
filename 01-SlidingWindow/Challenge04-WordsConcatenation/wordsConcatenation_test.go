package Challenge04_WordsConcatenation

import (
	"fmt"
	"testing"
)

func TestWordConcatenation(t *testing.T) {
	testCases := []struct {
		input    string
		words    []string
		expected []int
	}{
		{"catfoxcat", []string{"cat", "fox"}, []int{0, 3}},
		{"catcatfoxfox", []string{"cat", "fox"}, []int{3}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v, pattern: %v", tc.input, tc.words), func(t *testing.T) {
			result := wordConcatenation(tc.input, tc.words)

			for i, v := range tc.expected {
				if v != result[i] {
					t.Errorf("Index %d; exp %d, got %d", i, v, result[i])
				}
			}
		})
	}
}

/*func TestWordConcatenation2(t *testing.T) {
	testCases := []struct {
		input    string
		words    []string
		expected []int
	}{
		{"catfoxcat", []string{"cat", "fox"}, []int{0, 3}},
		{"catcatfoxfox", []string{"cat", "fox"}, []int{3}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("input: %v, pattern: %v", tc.input, tc.words), func(t *testing.T) {
			result := wordConcatenation2(tc.input, tc.words)

			for i, v := range tc.expected {
				if v != result[i] {
					t.Errorf("Index %d; exp %d, got %d", i, v, result[i])
				}
			}
		})
	}
}*/

func BenchmarkWordConcatenation(b *testing.B) {
	b.ResetTimer()

	str := "catcatfoxfox"
	words := []string{"cat", "fox"}

	for i := 0; i < b.N; i++ {
		_ = wordConcatenation(str, words)
	}
}

/*func BenchmarkWordConcatenation2(b *testing.B) {
	b.ResetTimer()

	str := "catcatfoxfox"
	words := []string{"cat", "fox"}

	for i := 0; i < b.N; i++ {
		_ = wordConcatenation2(str, words)
	}
}
*/
