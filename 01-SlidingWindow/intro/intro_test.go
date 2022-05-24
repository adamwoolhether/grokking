package intro

import "testing"

func TestFindAveragesOfSubarrays(t *testing.T) {
	arr := []float64{1, 3, 2, 6, -1, 4, 1, 8, 2}
	K := 5
	expected := []float64{2.2, 2.8, 2.4, 3.6, 2.8}
	expectedLen := len(expected)

	result := findAveragesOfSubarrays(K, arr)
	resultLen := len(result)

	if expectedLen != resultLen {
		t.Fatalf("Exp result length %d, got %d", expectedLen, resultLen)
	}

	for i := 0; i < expectedLen; i++ {
		if expected[i] != result[i] {
			t.Errorf("Index %d; exp %2f, got %2f", i, expected[i], result[i])
		}
	}
}

func TestFindAveragesOfSubarraysBruteForce(t *testing.T) {
	arr := []float64{1, 3, 2, 6, -1, 4, 1, 8, 2}
	K := 5
	expected := []float64{2.2, 2.8, 2.4, 3.6, 2.8}
	expectedLen := len(expected)

	result := findAveragesOfSubarraysBruteForce(K, arr)
	resultLen := len(result)

	if expectedLen != resultLen {
		t.Fatalf("Exp result length %d, got %d", expectedLen, resultLen)
	}

	for i := 0; i < expectedLen; i++ {
		if expected[i] != result[i] {
			t.Errorf("Index %d; exp %2f, got %2f", i, expected[i], result[i])
		}
	}
}

func BenchmarkFindAveragesOfSubarrays(b *testing.B) {
	b.ResetTimer()

	arr := []float64{1, 3, 2, 6, -1, 4, 1, 8, 2}
	k := 5

	for i := 0; i < b.N; i++ {
		_ = findAveragesOfSubarrays(k, arr)
	}
}

func BenchmarkFindAveragesOfSubarraysBruteForce(b *testing.B) {
	b.ResetTimer()

	arr := []float64{1, 3, 2, 6, -1, 4, 1, 8, 2}
	k := 5

	for i := 0; i < b.N; i++ {
		_ = findAveragesOfSubarraysBruteForce(k, arr)
	}
}