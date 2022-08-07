package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type JudgeFunc func(any) bool

func MultiJudge(value any, judgeFuncs []JudgeFunc) bool {
	var res int32
	wg := sync.WaitGroup{}

	for _, judgeFunc := range judgeFuncs {
		wg.Add(1)
		go func(judgeFunc JudgeFunc) {
			defer wg.Done()
			if atomic.LoadInt32(&res) == 0 {
				if judgeFunc(value) {
					atomic.StoreInt32(&res, 1)
				}
			}
		}(judgeFunc)
	}
	wg.Wait()
	return res != 0
}

func main() {

	res := MultiJudge(4, []JudgeFunc{
		func(a any) bool {
			fmt.Println("judge 1 called")
			return a.(int) == 1
		},
		func(a any) bool {
			fmt.Println("judge 2 called")
			return a.(int) == 2
		},
		func(a any) bool {
			fmt.Println("judge 3 called")
			return a.(int) == 3
		},
		func(a any) bool {
			fmt.Println("judge 4 called")
			return a.(int) == 4
		},
	})

	fmt.Println("MultiJudge res", res)
}
