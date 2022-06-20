package Challenge02_RearrangeALinkedList

import (
	"testing"
)

func TestReorder(t *testing.T) {
	head0 := &Node[int]{value: 2}
	head0.next = &Node[int]{value: 4}
	head0.next.next = &Node[int]{value: 6}
	head0.next.next.next = &Node[int]{value: 8}
	head0.next.next.next.next = &Node[int]{value: 10}
	head0.next.next.next.next.next = &Node[int]{value: 12}

	head1 := &Node[int]{value: 2}
	head1.next = &Node[int]{value: 4}
	head1.next.next = &Node[int]{value: 6}
	head1.next.next.next = &Node[int]{value: 8}
	head1.next.next.next.next = &Node[int]{value: 10}

	testCases := []struct {
		name string
		head *Node[int]
		exp  []int
	}{
		{"head0", head0, []int{2, 12, 4, 10, 6, 8}},
		{"head1", head1, []int{2, 10, 4, 8, 6}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reorder(tc.head)

			tempArray := make([]int, 0, len(tc.exp))
			temp := tc.head
			for temp != nil {
				tempArray = append(tempArray, temp.value)
				temp = temp.next
			}

			for i := range tc.exp {
				if tempArray[i] != tc.exp[i] {
					t.Errorf("exp value %d, got %d", tc.exp[i], tempArray[i])
				}
			}
		})
	}
}
