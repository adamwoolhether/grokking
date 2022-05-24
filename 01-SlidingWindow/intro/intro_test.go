package intro

import "testing"

// type fnSig func[T numConstraints](k int, arr []T) []T

func TestFindAveragesOfSubarrays(t *testing.T) {
	arr := []float64{1, 3, 2, 6, -1, 4, 1, 8, 2}
	K := 5
	expected := []float64{2.2, 2.8, 2.4, 3.6, 2.8}
	expectedLen := len(expected)

	result := findAveragesOfSubarraysSlidingWindow(K, arr)
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

// func testAverageOfSubarrays(t *testing.T, ) {
//
// }
