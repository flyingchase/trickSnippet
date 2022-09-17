package base

import (
	"fmt"
	"unsafe"
)

func RangeCopy() {
	var v []int = []int{1, 2, 3}
	// 不会无限循环，v=[]int{1,2,3,1,2,3}
	// 首次时拷贝副本的 len 一次
	for _, value := range v {
		v = append(v, value)
	}
	fmt.Println(v)
}

func UnsafePointer() {
	i := 10
	ip := &i
	var fp *float64 = (*float64)(unsafe.Pointer(ip))

	*fp *= 3
	fmt.Println(fp)
	fmt.Println(&i)
	fmt.Println(i)
	fmt.Println("________________")

	u := new(user)
	fmt.Println(*u)

	pName := (*string)(unsafe.Pointer(u))
	*pName = "anyOne"
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u)) + unsafe.Offsetof(u.age)))
	*pAge = 10086

	fmt.Println(*u)
	fmt.Println(u)

}

type user struct {
	name string
	age  int
}
