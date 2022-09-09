package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// MergeErrors merges multiple channles of errors
func MergeErrors(cs ...<-chan error) <-chan error {
	var wg sync.WaitGroup
	out := make(chan error, len(cs))

	outPut := func(c <-chan error) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go outPut(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// WaitForPipeline waits for results from all error goroutines are done.
// return early on the first error
func WaitForPipeline(err ...<-chan error) error {
	errChan := MergeErrors(err...)
	for err := range errChan {
		if err != nil {
			return err
		}
	}
	return nil
}

func minimalPipelineStage(ctx context.Context) (<-chan error, error) {
	errChan := make(chan error, 1)
	go func() {
		defer close(errChan)
		// do sth useful here
	}()

	return errChan, nil

}

func lineListSource(ctx context.Context, lines ...string) (<-chan string, <-chan error, error) {
	if len(lines) == 0 {
		return nil, nil, fmt.Errorf("no lines provided")
	}
	out, errChan := make(chan string), make(chan error, 1)
	go func() {
		defer close(out)
		defer close(errChan)

		for lineIndex, line := range lines {
			if line == "" {
				errChan <- fmt.Errorf("line %v is empty", lineIndex+1)
			}

			// Send the data to the out channel but return early if the ctx has been cancelled.
			select {
			case out <- line:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out, errChan, nil
}

func lineParser(ctx context.Context, base int, in <-chan string) (<-chan int64, <-chan error, error) {
	if base < 2 {
		return nil, nil, fmt.Errorf("invalid base %v", base)
	}

	out, errChan := make(chan int64), make(chan error, 1)
	go func() {
		defer close(out)
		defer close(errChan)

		for line := range in {
			n, err := strconv.ParseInt(line, base, 64)
			if err != nil {
				errChan <- err
				return
			}

			// Send the data to the outPut channel but return early if the ctx has been cancelled.
			select {
			case out <- n:
			case <-ctx.Done():
				return
			}
		}
	}()

	return out, errChan, nil
}

func sink(ctx context.Context, in <-chan int64) (<-chan error, error) {

	errChan := make(chan error, 1)
	go func() {
		defer close(errChan)
		for n := range in {
			if n > 100 {
				errChan <- fmt.Errorf("num %v is too large", n)
				return
			}
			fmt.Printf("sink: %v\n", n)
		}
	}()
	return errChan, nil
}

func runSimplePipeline(base int, lines []string) error {
	fmt.Printf("runSimplePipeline: base=%v, lines=%v\n", base, lines)

	ctx, cancelFunc := context.WithCancel(context.Background())

	defer cancelFunc()

	var errChanList []<-chan error

	// Source pipeline stage.
	lineChan, errChan, err := lineListSource(ctx, lines...)
	if err != nil {
		return err
	}
	errChanList = append(errChanList, errChan)

	// Transformer pipeline stage.
	numChan, errChan, err := lineParser(ctx, base, lineChan)
	if err != nil {
		return err
	}
	errChanList = append(errChanList, errChan)

	// Srink pipeline stage.
	errChan, err = sink(ctx, numChan)
	if err != nil {
		return err
	}
	errChanList = append(errChanList, errChan)
	fmt.Println("Pipeline started. Waitting for pipeline to complete.")
	return WaitForPipeline(errChanList...)
}

func randomNumberSource(ctx context.Context, seed int64) (<-chan string, <-chan error, error) {
	out, errChan := make(chan string), make(chan error, 1)
	random := rand.New(rand.NewSource(seed))
	go func() {
		defer close(out)
		defer close(errChan)

		for {
			n := random.Intn(100)
			line := fmt.Sprintf("%v", n)
			select {
			case out <- line:
			case <-ctx.Done():
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	return out, errChan, nil
}

func runPipelineWithTimeout() error {
	fmt.Println("runPipelineWithTimeout")
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	var errChanList []<-chan error
	// Source pipeline stage.
	linec, errc, err := randomNumberSource(ctx, 3)
	if err != nil {
		return err
	}
	errChanList = append(errChanList, errc)

	// Transformer pipeline stage.
	numberc, errc, err := lineParser(ctx, 10, linec)
	if err != nil {
		return err
	}
	errChanList = append(errChanList, errc)

	// Sink pipeline stage.
	errc, err = sink(ctx, numberc)
	if err != nil {
		return err
	}
	errChanList = append(errChanList, errc)

	fmt.Println("Pipeline started. Waiting for pipeline to complete.")

	// Start a goroutine that will cancel this pipeline in 10 seconds.
	go func() {
		time.Sleep(10 * time.Second)
		fmt.Println("Cancelling context.")
		cancelFunc()
	}()

	return WaitForPipeline(errChanList...)

}
func main() {
	if err := runSimplePipeline(10, []string{"3", "2", "1"}); err != nil {
		fmt.Println(err)
	}
	// if err := runSimplePipeline(1, []string{"3", "2", "1"}); err != nil {
	// 	fmt.Println(err)
	// }
	// if err := runSimplePipeline(2, []string{"1010", "1100", "1000"}); err != nil {
	// 	fmt.Println(err)
	// }
	// if err := runSimplePipeline(2, []string{"1010", "1100", "2000", "1111"}); err != nil {
	// 	fmt.Println(err)
	// }
	// if err := runSimplePipeline(10, []string{"1", "10", "100", "1000"}); err != nil {
	// 	fmt.Println(err)
	// }
	//
	// if err := runPipelineWithTimeout(); err != nil {
	// 	fmt.Println(err)
	// }
}
