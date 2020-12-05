package sort

func SortColors(nums []int)  {
	//边界问题 如果数组大小小于等于1，直接返回
	if len(nums) <= 1 {
		return
	}
	quickSort(nums,0,len(nums) - 1)
}

func quickSort(nums[] int,left,right int)  {
	//边界问题
	if right > left {
		pivotIndex := (right - left + 1) / 2 + left //求余取数问题
		pivotIndex = partition(nums, left, right, pivotIndex)
		quickSort(nums,left,pivotIndex-1)
		quickSort(nums, pivotIndex +1, right)
	}

}

func partition(nums[] int, left,right,pivot int) int  {

	pivotValue := nums[pivot]
	swap(nums,pivot,right)
	storeIndex := left //递归定义边界问题
	for i := left; i < right; i++ {
		if nums[i] <= pivotValue{
			swap(nums,storeIndex,i)
			storeIndex++
		}
	}
	swap(nums,storeIndex,right)
	return storeIndex
}

func swap(nums[] int, x,y int)  {
	origin := nums[x]
	nums[x] = nums[y]
	nums[y] = origin
}