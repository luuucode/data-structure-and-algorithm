package linklist

import "testing"

func TestCount(t *testing.T) {

	/*node4 := ListNode{1, &node5}
	node3 := ListNode{6, &node4}
	node2 := ListNode{1, &node3}*/
	node1 := ListNode{2, nil}
	head := ListNode{1, &node1}

	t.Log(Count(&head))
	t.Log(IsPalindrome(&head))
}