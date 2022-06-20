package Challenge02_RearrangeALinkedList

import "golang.org/x/exp/constraints"

/*
Problem Statement

Given the head of a Singly LinkedList, write a method to modify the LinkedList
such that the nodes from the second half of the LinkedList are inserted alternately
to the nodes from the first half in reverse order. So if the LinkedList has nodes
1 -> 2 -> 3 -> 4 -> 5 -> 6 -> null, your method should return 1 -> 6 -> 2 -> 5 -> 3 -> 4 -> null.

Your algorithm should not use any extra space and the input LinkedList should be modified in-place.

Example 1:
Input: 2 -> 4 -> 6 -> 8 -> 10 -> 12 -> null
Output: 2 -> 12 -> 4 -> 10 -> 6 -> 8 -> null

Example 2:
Input: 2 -> 4 -> 6 -> 8 -> 10 -> null
Output: 2 -> 10 -> 4 -> 8 -> 6 -> null

Solution

This problem shares similarities with Palindrome LinkedList.
To rearrange the given LinkedList we will follow the following steps:

1. 	We can use the Fast & Slow pointers method similar to Middle
	of the LinkedList to find the middle node of the LinkedList.
2. 	Once we have the middle of the LinkedList, we will reverse the second half of the LinkedList.
3. 	Finally, we’ll iterate through the first half and the reversed second half to produce a
	LinkedList in the required order.
*/

type numbers interface {
	constraints.Float | constraints.Integer
}

type Node[T numbers] struct {
	value T
	next  *Node[T]
}

/*
Time Complexity: O(N) where 'N' is the number of nodes in the linked list.
Space Complexity: Constant space, O(1)
*/
func reorder[T numbers](head *Node[T]) {
	if head == nil || head.next == nil {
		return
	}

	// find the middle of the linkedList (same code from `middleOfTheLinkedList')
	slow, fast := head, head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	// 'slow' points to the middle node, reverse it to get the second half.
	headSecondHalf := reverse(slow)
	headFirstHalf := head

	// rearrange to produce the linkedlist in the required order.
	for headFirstHalf != nil && headSecondHalf != nil {
		temp := headFirstHalf.next
		headFirstHalf.next = headSecondHalf
		headFirstHalf = temp

		temp = headSecondHalf.next
		headSecondHalf.next = headFirstHalf
		headSecondHalf = temp
	}

	// set the next of the last node to nil
	if headFirstHalf != nil {
		headFirstHalf.next = nil
	}

}

func reverse[T numbers](head *Node[T]) *Node[T] {
	var prev *Node[T]

	for head != nil {
		next := head.next
		head.next = prev
		prev = head
		head = next
	}

	return prev
}
