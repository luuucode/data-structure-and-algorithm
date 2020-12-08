package array

import "fmt"

func Arrays()  {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0],a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 4, 5, 6}
	fmt.Println(primes)
}

func Add(arr[] int,index int,a int) []int {
	ls := len(arr)
	if index <= 0 || index > ls {
		fmt.Println("error")
		return []int{}
	}
	res := arr[:ls]
	res = append(res, 0)
	for i := len(res) -1; i > index ; i-- {
		res[i] = res[i-1]
	}
	res[index] = a
	return res
}