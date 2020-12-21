package linklist

/**
Reverse a linked list from position m to n. Do it in one-pass.

Note:1 ≤ m ≤ n ≤ length of list.

Example:

Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//TODO
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	//判断边界问题
	if head == nil || m == n {
		return head
	}
	return head
}