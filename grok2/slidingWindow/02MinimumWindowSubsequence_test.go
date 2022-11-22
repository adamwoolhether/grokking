package slidingWindow

import "testing"

func Test_minWindow(t *testing.T) {
	tests := []struct {
		str1 string
		str2 string
		want string
	}{
		{"abcdebdde", "bde", "bcde"},
		{"fgrqsqsnodwmxzkzxwqegkndaa", "kzed", "kzxwqegknd"},
		{"michmznaitnjdnjkdsnmichmznait", "michmznait", "michmznait"},
		{"afgegrwgwga", "aa", "afgegrwgwga"},
		{"abababa", "ba", "ba"},
		{"Hello how are you?", "ok", ""},
	}
	for _, tt := range tests {
		t.Run(tt.str1, func(t *testing.T) {
			if got := minWindow(tt.str1, tt.str2); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
