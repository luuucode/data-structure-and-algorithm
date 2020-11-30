package swap

//交换函数
func Swap(x,y int, arr[] int) {
	var origin = arr[x]
	arr[x] = arr[y]
	arr[y] = origin
}

func Min(x,y int,arr[] int) (r int,i int) {
	if arr[x] > arr[y] {
		return arr[y],y
	}
	return arr[x],x
}