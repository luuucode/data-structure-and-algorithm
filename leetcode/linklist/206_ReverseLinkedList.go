package linklist

/**
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//当前版本 时间复杂度O(n) 空间复杂度O(n)
func ReverseList1(head *ListNode) *ListNode {
	res := &ListNode{head.Val,nil}
	node := head.Next
	for node != nil {
		res = &ListNode{node.Val,res}
		node = node.Next
	}
	return res
}

//双指针做法（我想的第二个思路）
func ReverseList2(head *ListNode) *ListNode{
	var cur *ListNode = nil
	//边界问题
	if head == nil{
		return head
	}
	prev := head
	for prev != nil {
		next := prev.Next
		prev.Next = cur
		cur = prev
		prev = next
	}
	return cur
}
//妖魔化指针法（我想的第三个思路）
func ReverseList3(head *ListNode) *ListNode{
	//边界问题
	if head == nil {
		return head
	}
	//cur 作用：指向反转结果的头节点
	cur := head
	for head.Next != nil {
		temp := head.Next
		head.Next = temp.Next
		temp.Next = cur
		cur = temp
	}
	return cur
}