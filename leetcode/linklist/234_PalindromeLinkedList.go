package linklist

/**
Given a singly linked list, determine if it is a palindrome.

Example 1:

Input: 1->2
Output: false
Example 2:

Input: 1->2->2->1
Output: true
Follow up:
Could you do it in O(n) time and O(1) space?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type ListNode struct {
		Val int
		Next *ListNode
}
//TODO 优化 空间复杂度为O(1)
func IsPalindrome(head *ListNode) bool {
	count := Count(head)
	if count == 1{
		return true
	}
	half := count/2
	sign := 0
	m := make(map[int]int, 0)
	node := head
	for node != nil {
		if sign < half {
			m[sign] = node.Val
		} else if count%2 == 1 && sign == half{
			sign += 1
			node = node.Next
			continue
		}else {
			if m[count - sign - 1] != node.Val{
				return false;
			}
		}
		sign += 1
		node = node.Next
	}
	return true
}


func Count(head *ListNode) int {
	sum := 0
	node := head
	for node != nil {
		sum += 1
		node = node.Next
	}
	return sum
}