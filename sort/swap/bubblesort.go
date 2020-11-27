package swap

func BubbleSort(arr[] int,len int)  {
	//递增排序
	for i := 0; i < len; i++ {
		Float(arr,len)
	}
}

func Float(arr[] int,len int)  {
	//上浮
	for i := 0; i < len-1; i++ {
		if arr[i] < arr[i+1]  {
			Swap(i, i + 1, arr)
		}
	}
}