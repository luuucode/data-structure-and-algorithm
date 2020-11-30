package insert

import (
	"data-structure-and-algorithm/sort/swap"
)

/**
   插入排序（从小到大）
   从头到尾进行标记，对第i个数跟前i-1数进行比较，如果比前面数小，就进行交换
 */
func InsertionSort(arr[] int, left int, right int)  {
	for i := left; i <= right; i++ {
		for j := i; j > left ; j-- {
			if arr[j] < arr[j-1] {
				swap.Swap(j,j-1,arr)
			}
		}
	}
}