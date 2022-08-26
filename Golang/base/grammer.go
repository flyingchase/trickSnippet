package base

import (
	"fmt"
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
