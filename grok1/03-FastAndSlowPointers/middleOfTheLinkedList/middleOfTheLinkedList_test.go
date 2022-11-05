package middleOfTheLinkedList

import "testing"

func TestFindMiddle(t *testing.T) {
	head0 := &Node[int]{value: 1}
	head0.next = &Node[int]{value: 2}
	head0.next.next = &Node[int]{value: 3}
	head0.next.next.next = &Node[int]{value: 4}
	head0.next.next.next.next = &Node[int]{value: 5}

	exp := head0.next.next
	if result := findMiddle(head0); exp != result {
		t.Errorf("exp %v, got %v", exp, result)
	}

	head1 := head0
	head1.next.next.next.next.next = &Node[int]{value: 6}

	exp = head0.next.next.next
	if result := findMiddle(head1); exp != result {
		t.Errorf("exp %v, got %v", exp, result)
	}

	head2 := head1
	head2.next.next.next.next.next.next = &Node[int]{value: 7}

	// exp same as above, as len is not even.
	if result := findMiddle(head2); exp != result {
		t.Errorf("exp %v, got %v", exp, result)
	}

}
