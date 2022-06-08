package startOfLinkedListCycle

import "golang.org/x/exp/constraints"

/*
Problem Statement:
Given the head of a Singly LinkedList that contains a cycle, write a function to find the starting node of the cycle.

Solution:
If we know the length of the LinkedList cycle, we can find the start of the cycle through the following steps:

1. Take two pointers. Let’s call them pointer1 and pointer2.
2. Initialize both pointers to point to the start of the LinkedList.
3. We can find the length of the LinkedList cycle using the approach discussed in  LinkedList Cycle.
	Let’s assume that the length of the cycle is ‘K’ nodes.
4. Move pointer2 ahead by ‘K’ nodes.
5. Now, keep incrementing pointer1 and pointer2 until they both meet.
6. As pointer2 is ‘K’ nodes ahead of pointer1, which means,
	pointer2 must have completed one loop in the cycle when both pointers meet.
	Their meeting point will be the start of the cycle.

Let’s visually see this with the above-mentioned Example-1: (https://designgurus.org/path-player?courseid=grokking-the-coding-interview&unit=grokking-the-coding-interview_1628743554403_14Unit)

Use the algo in startOfLinkedListCycle to get the length.
*/

type numbers interface {
	constraints.Float | constraints.Integer
}

type Node[T numbers] struct {
	value T
	next  *Node[T]
}

func findCycleStart[T numbers](head *Node[T]) *Node[T] {
	cycleLen := 0
	slow, fast := head, head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next

		if slow == fast {
			cycleLen = cycleLength(slow)
			break
		}
	}

	return findStart(head, cycleLen)
}

func cycleLength[T numbers](slow *Node[T]) int {
	current := slow
	cycleLen := 0

	for {
		current = current.next
		cycleLen++

		if current == slow {
			break
		}
	}

	return cycleLen
}

func findStart[T numbers](head *Node[T], cycleLen int) *Node[T] {
	pointer1, pointer2 := head, head

	for cycleLen > 0 {
		pointer2 = pointer2.next
		cycleLen--
	}

	// Increment both pointers until they meet at the start.
	for pointer1 != pointer2 {
		pointer1 = pointer1.next
		pointer2 = pointer2.next
	}

	return pointer1
}