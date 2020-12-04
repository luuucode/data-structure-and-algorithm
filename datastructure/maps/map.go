package maps

import (
	"data-structure-and-algorithm/datastructure/array"
	"fmt"
)

var m map[string]array.Vertex

func Map() {
	//定义方式1
	m = make(map[string]array.Vertex)
	m["Bell Labs"] = array.Vertex{
		40,
		34,
	}
	fmt.Println(m)
	for s := range m {
		fmt.Println(s)
	}
	for s, vertex := range m {
		fmt.Println(s,":",vertex)
	}
	//定义方式2
	m2 := map[string]array.Vertex{
		"张三":{2,3},
		"李四":{45,43},
	}
	fmt.Println(m2)
	for s, vertex := range m2 {
		fmt.Println(s,":",vertex)
	}

	//赋值&取值
	m3 := make(map[string]int)
	m3["a"] = 10
	m3["b"] = 11
	fmt.Println("a value :",m3["a"])

	//删除
	delete(m3,"a") //delete(m,key)

	//判断是否存在
	v,ok := m3["a"]
	fmt.Println("a value :",v,"present?",ok) // value 值是0  ok = false

}
