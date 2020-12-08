package sort

/**
Given a list of non-negative integers nums, arrange them such that they form the largest number.

Note: The result may be very large, so you need to return a string instead of an integer.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/largest-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func LargeNumber(nums[] int) string {
	return "999"
}

func getMaxHighNum(num int) int {
	m := num
	divide := pow(10, size(num)-1)
	if divide > 0{
		m = num / divide;
	}
	return m
}

func size(num int) int {
	r := 0
	for num > 0 {
		num /= 10
		r++
	}
	return r
}

func pow(o int,n int) int {
	res := o
	for i := 0; i < n - 1; i++ {
		res *= o
	}
	return res
}