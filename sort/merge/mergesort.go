package merge

/**
 递归法 top down
 */
func MergeSortTopDown(arr[] int) (r[] int){
	//split
	size := len(arr)
	if size <= 1 {
		return arr
	}
	mid := size / 2
	left := arr[0:mid]
	right := arr[mid:]
	leftSize := len(left)
	if leftSize >= 2 {
		left = MergeSortTopDown(left)
	}
	rightSize := len(right)
	if rightSize >= 2 {
		right = MergeSortTopDown(right)
	}
	//merge
	j := 0
	t := 0
	//分配O(N)空间数组来存放结果
	list := make([]int,size)
	for i := 0; i < size; i++ {
		if j < leftSize && t <rightSize {
			if left[j] <= right[t] {
				list[i] = left[j]
				j++
			}else {
				list[i] = right[t]
				t++
			}
		} else if j >= leftSize {
			list[i] = right[t]
			t++
		}else if t >= rightSize {
			list[i] = left[j]
			j++
		}
	}
	return list
}

