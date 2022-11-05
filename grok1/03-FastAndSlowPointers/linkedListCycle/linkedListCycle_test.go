package linkedListCycle

import "testing"

func TestHasCycle(t *testing.T) {
	head0 := &Node[int]{value: 1}
	head0.next = &Node[int]{value: 2}
	head0.next.next = &Node[int]{value: 3}
	head0.next.next.next = &Node[int]{value: 4}
	head0.next.next.next.next = &Node[int]{value: 5}
	head0.next.next.next.next.next = &Node[int]{value: 6}

	exp := false
	if result := hasCycle(head0); result != exp {
		t.Errorf("exp %t, got %t", exp, result)
	}

	head1 := head0
	head1.next.next.next.next.next.next = head1.next.next

	exp = true
	if result := hasCycle(head1); result != exp {
		t.Errorf("exp %t, got %t", exp, result)
	}

	head2 := head1
	head1.next.next.next.next.next.next = head1.next.next.next

	exp = true
	if result := hasCycle(head2); result != exp {
		t.Errorf("exp %t, got %t", exp, result)
	}
}

func TestFindCycleLength(t *testing.T) {
	head0 := &Node[int]{value: 1}
	head0.next = &Node[int]{value: 2}
	head0.next.next = &Node[int]{value: 3}
	head0.next.next.next = &Node[int]{value: 4}
	head0.next.next.next.next = &Node[int]{value: 5}
	head0.next.next.next.next.next = &Node[int]{value: 6}

	exp := 0
	if result := findCycleLength(head0); result != exp {
		t.Errorf("exp %v, got %v", exp, result)
	}

	head1 := head0
	head1.next.next.next.next.next.next = head1.next.next

	exp = 4
	if result := findCycleLength(head1); result != exp {
		t.Errorf("exp %v, got %v", exp, result)
	}

	head2 := head1
	head1.next.next.next.next.next.next = head1.next.next.next

	exp = 3
	if result := findCycleLength(head2); result != exp {
		t.Errorf("exp %v, got %v", exp, result)
	}
}
