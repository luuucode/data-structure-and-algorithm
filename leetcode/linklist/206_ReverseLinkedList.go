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

//todo
func ReverseList2(head *ListNode) *ListNode{
	var prev *ListNode = nil
	for head.Next != nil {

	}
	return prev
}