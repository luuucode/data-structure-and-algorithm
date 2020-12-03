package array

import "fmt"

type Vertex struct {
	X int
	Y int
}

//struct 初始化定义
var (
	v1 = Vertex{1,2} //has type Vertex
	v2 = Vertex{X: 1} //Y：0 默认
	v3 = Vertex{} //默认X:0 Y:0
	p = &Vertex{1,2} //has type *Vertex

)


func Structs()  {
	v := Vertex{1,2}
	v.X = v.X * 2
	v.Y = v.Y * 3
	fmt.Println(v)
}

func PointerStructs()  {
	v := Vertex{1,2}
	p := &v
	fmt.Println(p) // point address
	p.X = 1e9
	fmt.Println(*p)
	fmt.Println(v)
	(*p).X = 2e9
	fmt.Println(v)
}

func InitStructs()  {
	fmt.Println(v1,v2,v3,p)
}