package single_filght

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var count int32

func SingleFilghtEffect() {
	time.AfterFunc(1*time.Second, func() {
		atomic.AddInt32(&count, -count)
	})

	var (
		wg  sync.WaitGroup
		now = time.Now()
		n   = 10000000
		sg  = &singleflight.Group{}
	)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			// mockFunc(i)
			singleflightMockFunc(sg, i)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("request = %d, waste time = %s", n, time.Since(now))
}

func mockFunc(i int) (int, error) {
	return rand.Intn(i), nil
}

func singleflightMockFunc(sg *singleflight.Group, id int) (int, error) {
	v, err, _ := sg.Do(fmt.Sprintf("%s", id), func() (interface{}, error) {
		return mockFunc(id)
	})
	return v.(int), err

}
