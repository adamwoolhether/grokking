package slidingWindow

import "sync"

/*
Statement:
Given an integer array and a window of size `windowSize`,find the current
maximum value in the window as it slides through the entire array.
Note: If the window size is greater than the array size, we consider the entire array as a single window.

Constraints:
- 1 <= nums.length <= 10^3
- -10^4 <= nums[i] <=10^4
- 1 <= windowSize

Steps:
1. Process first k elements to initiate deque
2. Iterate array
3. In deque, keep indexes of elements of the current sliding window.
4. Remove indexes of all elements smaller than current element from deque
5. Add current element to deque.
6. Add first element of deque to output list
7. Repeat above until entire array is traversed.
*/

// findMaxSlidingWindow return an array of the max size int for each window in the given input.
func findMaxSlidingWindow(nums []int, windowSize int) []int {
	if len(nums) == 0 {
		return nil
	}
	if windowSize > len(nums) {
		windowSize = len(nums)
	}

	result := make([]int, len(nums)-windowSize+1)
	dequeue := NewDequeue[int]()

	for i := 0; i < len(nums); i++ {
		// Remove indexes of elements that are smaller than the current element.
		for dequeue.size() > 0 && nums[i] >= nums[dequeue.peakBack()] {
			dequeue.dequeueBack()
		}

		// Remove first index from deque if its no longer in current window.
		if dequeue.size() > 0 && dequeue.peak() <= i-windowSize {
			dequeue.dequeue()
		}

		// Add current element index to back of queue.
		dequeue.enqueue(i)

		if i >= windowSize-1 {
			// result = append(result, nums[dequeue.peak()])
			result[i-windowSize+1] = nums[dequeue.peak()]
		}
	}

	return result
}

// /////////////////////////////////////////////////////////////////

type Dequeue[T any] struct {
	mu   sync.RWMutex
	data []T
}

func NewDequeue[T any]() *Dequeue[T] {
	return &Dequeue[T]{
		mu:   sync.RWMutex{},
		data: []T{},
	}
}

func (q *Dequeue[T]) enqueue(v T) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.data = append(q.data, v)
}

func (q *Dequeue[T]) dequeue() T {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.isEmpty() {
		return q.nilType()
	}

	v := q.data[0]
	// q.data[0] = nil
	q.data = q.data[1:]

	return v
}

func (q *Dequeue[T]) dequeueBack() T {
	q.mu.Lock()
	defer q.mu.Unlock()

	if q.isEmpty() {
		return q.nilType()
	}

	v := q.data[len(q.data)-1]
	q.data = q.data[:len(q.data)-1]

	return v
}

func (q *Dequeue[T]) peak() T {
	q.mu.RLock()
	defer q.mu.RUnlock()

	if q.isEmpty() {
		return q.nilType()
	}

	return q.data[0]
}

func (q *Dequeue[T]) peakBack() T {
	q.mu.RLock()
	defer q.mu.RUnlock()

	if q.isEmpty() {
		return q.nilType()
	}

	return q.data[len(q.data)-1]
}

func (q *Dequeue[T]) size() int {
	q.mu.RLock()
	defer q.mu.RUnlock()

	return len(q.data)
}

func (q *Dequeue[T]) isEmpty() bool {
	// no locking, as this is internal and should be done by caller.
	return len(q.data) == 0
}

func (q *Dequeue[T]) nilType() T {
	var t T

	return t
}
