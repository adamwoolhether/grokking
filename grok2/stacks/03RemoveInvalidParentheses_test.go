package stacks

import (
	"testing"
)

func TestMinRemoveParentheses(t *testing.T) {
	testCases := []struct {
		input string
		exp   string
	}{
		{"()", "()"},
		{"(", ""},
		{")", ""},
		{")(", ""},
		{"ab)ca(so)(sc(s)(", "abca(so)sc(s)"},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			if result := minRemoveParentheses(tc.input); result != tc.exp {
				t.Errorf("exp: %s, got: %s", tc.exp, result)
			}
		})
	}
}

func BenchmarkMinRemoveParentheses(b *testing.B) {
	str := "ab)ca(so)(sc(s)("
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = minRemoveParentheses(str)
	}
}

func BenchmarkMinRemoveParenthesesWithMap(b *testing.B) {
	str := "ab)ca(so)(sc(s)("
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = minRemoveParenthesesWithMap(str)
	}
}
