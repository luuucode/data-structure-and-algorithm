package sort

import (
	"math/rand"
)

func MergeIntervals(arr[][] int) (result[][] int) {

	//对2维数组按第一列数排序
	QuickSort2DArrByRow(arr,0,len(arr) - 1)
	r := make([][]int, 0)
	r = append(r,arr[0])
	for i := 1; i < len(arr); i++ {
		if arr[i][0] > r[len(r)-1][1]{
			r = append(r,arr[i])
		}else {
			//前一个数右极限大于后一个数左极限
			r[len(r)-1][1] = max(r[len(r)-1][1],arr[i][1])
		}
	}
	return r
}

func QuickSort2DArrByRow(arr[][] int,left,right int)  {
	if right > left {
		pivotIndex := rand.Intn(right -left + 1) + left
		pivotIndex = Partition2D(arr,left,right,pivotIndex)
		QuickSort2DArrByRow(arr,0,pivotIndex-1)
		QuickSort2DArrByRow(arr,pivotIndex +1,right)
	}


}

func Partition2D(a[][] int,left int,right int,pivotIndex int) (i int){
	startIndex := left
	pivotValue := a[pivotIndex][0]
	//移动标识位到最后
	Swap2D(a,pivotIndex,right)
	//从数组头开始比较
	for i := left; i < right; i++ {
		if a[startIndex][0] <= pivotValue{
			Swap2D(a,startIndex,i)
			startIndex++
		}
	}
	Swap2D(a,startIndex,right)
	return startIndex
}

func Swap2D(a[][] int, x,y int)  {
	swap := a[x]
	a[x] = a[y]
	a[y] = swap
}
func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}


