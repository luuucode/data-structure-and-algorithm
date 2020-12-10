package sort

func Intersection(nums1[] int,nums2[] int) []int {
	res := make([]int,0)
	m := convertCountMap(nums1)
	m2 := convertCountMap(nums2)
	for i := range m {
		v,ok := m2[i]
		if v >= 1 && ok == true{
			res = append(res, i)
		}
	}
	return res
}

/**
  map
	key is array number
	value is number count
 */
func convertCountMap(nums[] int) map[int]int {
	m1 := make(map[int]int)
	for i := range nums {
		v := nums[i]
		m1[v] ++
	}
	return m1
}