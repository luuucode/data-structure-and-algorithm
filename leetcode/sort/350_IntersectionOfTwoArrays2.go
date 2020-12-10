package sort

func Intersection2(nums1[] int,nums2[] int) []int {
	res := make([]int,0)
	m := make(map[int]int, 0)
	for _, v := range nums1 {
		m[v] ++
	}
	for _, v2 := range nums2 {
		if m[v2] > 0 {
			res = append(res, v2)
			m[v2]--
		}
	}
	return res
}