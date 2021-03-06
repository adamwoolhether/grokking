package middleOfTheLinkedList

import "golang.org/x/exp/constraints"

/*
Problem Statement:
Given the head of a Singly LinkedList, write a method to return the middle node of the LinkedList.

If the total number of nodes in the LinkedList is even, return the second middle node.

Example 1:
Input: 1 -> 2 -> 3 -> 4 -> 5 -> null
Output: 3

Example 2:
Input: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> null
Output: 4

Example 3:
Input: 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> null
Output: 4

Solution:
One brute force strategy could be to first count the number of nodes in the
LinkedList and then find the middle node in the second iteration.
Can we do this in one iteration?

We can use the Fast & Slow pointers method such that the fast pointer is
always twice the nodes ahead of the slow pointer. This way, when the fast pointer
reaches the end of the LinkedList, the slow pointer will be pointing at the middle node.
*/

type numbers interface {
	constraints.Float | constraints.Integer
}

type Node[T numbers] struct {
	value T
	next  *Node[T]
}

/*
Time Complexity: O(N), where 'N' is the number of nodes in the linked list.
Space Complexity: O(1)
*/

func findMiddle[T numbers](head *Node[T]) *Node[T] {
	slow, fast := head, head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}
