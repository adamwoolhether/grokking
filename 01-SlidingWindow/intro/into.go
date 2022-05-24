package intro

type numConstraints interface {
	int | uint | int32 | int64 |
	float32 | float64
}

func findAveragesOfSubarrays[T numConstraints](k int, arr []T) []T {
	result := make([]T, len(arr)-k+1)

	windowSum := T(0)
	windowStart := 0

	for windowEnd := 0; windowEnd < len(arr); windowEnd++ {
		windowSum += arr[windowEnd] // add the next element

		// Slide the window if we haven't hit window size of 'k'.
		if windowEnd >= k-1 {
			result[windowStart] = windowSum / T(k) // calculate the average
			windowSum -= arr[windowStart]          // subtract the element leaving the window
			windowStart++                          // slide the window ahead
		}
	}

	return result
}

func findAveragesOfSubarraysBruteForce[T numConstraints](k int, arr []T) []T {
	result := make([]T, len(arr)-k+1)

	for i := 0; i < len(arr)-k+1; i++ {
		// Find sum of next 'k' elements
		sum := T(0)

		for j := i; j < i+k; j++ {
			sum += arr[j]
		}

		result[i] = sum / T(k) // calculate the average
	}

	return result
}
