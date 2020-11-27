package swap

import (
	"math/rand"
)

func QuickSort(a[] int,left,right int)  {
	if right > left {
		pivotIndex := rand.Intn(right -left + 1) + left
		pivotNewIndex := Partition(a,left,right,pivotIndex)
		QuickSort(a,left,pivotNewIndex - 1)
		QuickSort(a,pivotNewIndex +1,right)
	}
}



func Partition(a[] int,left int,right int,pivotIndex int) (i int)  {
	pivotValue := a[pivotIndex] //获取pivot值
	Swap(pivotIndex,right,a)    //移动pivot到队尾
	storeIndex := left          //标识位 从数组最左开始
	//遍历数组n-1位数，如果当前值小于pivot值，发生交换
	//标识位与当前遍历索引位数值交换，标识位自增
	for i := left; i < right; i++ {
		if a[i] <= pivotValue {
			Swap(storeIndex,i,a)
			storeIndex++
		}
	}
	Swap(right,storeIndex,a)
	return storeIndex
}
