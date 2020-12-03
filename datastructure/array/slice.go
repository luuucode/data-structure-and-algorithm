package array

import "fmt"

func Slice()  {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Println(len(primes))
	//slice取数 x:y 不包含最右侧数据
	l := primes[1:3]
	fmt.Println(l)
	//slice cap 取决于依赖数组的长度
	fmt.Printf("len=%d cap=%d %v\n", len(l), cap(l), l)
	t := primes[0:]
	fmt.Println(t)
	s := primes[:6]
	fmt.Println(s)

	//Slices are like references to arrays 修改引用数据会更改原数组
	names := [4]string{"John", "Paul", "George", "Ringo"}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a,b)
	b[0] = "XXX"
	fmt.Println(a,b)
	fmt.Println(names)

	//creating a slice with make
	x := make([]int, 5)
	printSlice("x",x)
	x = append(x,1,2,3)
	printSlice("x append",x)
	y := make([]int,5,10)
	printSlice("y",y)
	z := make([]int,0)
	z = append(z,1,2,3,4)
	for i := range z {
		fmt.Println(z[i])
	}
	for i, r := range z {
		fmt.Println(i,r)
	}
	for _, value := range z {
		fmt.Println(value)
	}




}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}