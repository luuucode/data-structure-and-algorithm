package swap

//交换函数
func Swap(x,y int, arr[] int) {
	var origin = arr[x]
	arr[x] = arr[y]
	arr[y] = origin
}