package startOfLinkedListCycle

import "testing"

func TestFindCycleStart(t *testing.T) {
	head0 := &Node[int]{value: 1}
	head0.next = &Node[int]{value: 2}
	head0.next.next = &Node[int]{value: 3}
	head0.next.next.next = &Node[int]{value: 4}
	head0.next.next.next.next = &Node[int]{value: 5}
	head0.next.next.next.next.next = &Node[int]{value: 6}
	head0.next.next.next.next.next.next = head0.next.next

	exp := head0.next.next
	if result := findCycleStart(head0); exp != result {
		t.Errorf("exp %v, got %v", exp, result)
	}

	head1 := head0
	head1.next.next.next.next.next.next = head1.next.next.next

	exp = head1.next.next.next
	if result := findCycleStart(head1); exp != result {
		t.Errorf("exp %v, got %v", exp, result)
	}

	head2 := head1
	head2.next.next.next.next.next.next = head2

	exp = head2
	if result := findCycleStart(head2); exp != result {
		t.Errorf("exp %v, got %v", exp, result)
	}
}
