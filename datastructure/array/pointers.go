package array

import "fmt"
func Pointers()  {
	//指针操作
	i,j := 42,2701
	p := &i		//point to i
	fmt.Println(p) //print address of i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p = &j		  //point to j
	*p = *p / 37  //divide j through the pointer
	fmt.Println(j) //set the new value of j
}