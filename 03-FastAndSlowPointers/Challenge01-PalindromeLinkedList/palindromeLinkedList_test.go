package Challenge01_PalindromeLinkedList

import "testing"

func TestIsPalindromeLinkedList(t *testing.T) {
	head0 := &Node[int]{value: 2}
	head0.next = &Node[int]{value: 4}
	head0.next.next = &Node[int]{value: 6}
	head0.next.next.next = &Node[int]{value: 4}
	head0.next.next.next.next = &Node[int]{value: 2}

	head1 := &Node[int]{value: 2}
	head1.next = &Node[int]{value: 4}
	head1.next.next = &Node[int]{value: 6}
	head1.next.next.next = &Node[int]{value: 4}
	head1.next.next.next.next = &Node[int]{value: 2}
	head1.next.next.next.next.next = &Node[int]{value: 2}

	testCases := []struct {
		name string
		head *Node[int]
		exp  bool
	}{
		{"case 0", head0, true},
		{"case 1", head1, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := isPalindromeLinkedList(tc.head); tc.exp != got {
				t.Errorf("exp %t, got %t", tc.exp, got)
			}
		})
	}
}
