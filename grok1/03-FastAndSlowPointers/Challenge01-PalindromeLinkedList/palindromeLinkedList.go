package Challenge01_PalindromeLinkedList

import "golang.org/x/exp/constraints"

/*
Problem Statement:

Given the head of a Singly LinkedList, write a method to check if the LinkedList
is a palindrome or not.

Your algorithm should use constant space and the input LinkedList should be in the
original form once the algorithm is finished. The algorithm should have O(N) time
complexity where ‘N’ is the number of nodes in the LinkedList.

Example 1:
Input: 2 -> 4 -> 6 -> 4 -> 2 -> null
Output: true

Example 2:
Input: 2 -> 4 -> 6 -> 4 -> 2 -> 2 -> null
Output: false

Solution:
As we know, a palindrome LinkedList will have nodes values that read the
same backward or forward. This means that if we divide the LinkedList into
two halves, the node values of the first half in the forward direction
should be similar to the node values of the second half in the backward
direction. As we have been given a Singly LinkedList, we can’t move in the
backward direction. To handle this, we will perform the following steps:
	1. 	We can use the Fast & Slow pointers method similar to Middle of the
		LinkedList to find the middle node of the LinkedList.
	2. 	Once we have the middle of the LinkedList, we will reverse the second half.
	3. 	Then, we will compare the first half with the reversed second half to see
		if the LinkedList represents a palindrome.
	4. 	Finally, we will reverse the second half of the LinkedList again to revert
		and bring the LinkedList back to its original form.
*/

type numbers interface {
	constraints.Float | constraints.Integer
}

type Node[T numbers] struct {
	value T
	next  *Node[T]
}

/*
Time Complexity: O(N) where 'N' is the number of nodes in the linked list
Space Complexity: O(1)
*/

func isPalindromeLinkedList[T numbers](head *Node[T]) bool {
	if head == nil || head.next == nil { // return early for this edge case
		return true
	}

	// find the middle of the linked list
	middle := findMiddle(head)

	headSecondHalf := reverse(middle)    // reverse the second half
	copyHeadSecondHalf := headSecondHalf // store the head of reversed part to revert back to it later

	// compare the first and second half
	for head != nil && headSecondHalf != nil {
		if head.value != headSecondHalf.value {
			break // not a palindrome
		}
		head = head.next
		headSecondHalf = headSecondHalf.next
	}

	reverse(copyHeadSecondHalf) // revert the reverse of the second half

	if head == nil || headSecondHalf == nil { // if both halves match
		return true
	}

	return false
}

func findMiddle[T numbers](head *Node[T]) *Node[T] {
	slow, fast := head, head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
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
