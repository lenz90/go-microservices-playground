package main

import (
	"fmt"
	"sync"
)

func Pool(nums []int, w int) int {
	jobs := make(chan int)
	out := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < w; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range jobs {
				out <- n * n
			}
		}()
	}
	go func() {
		for _, n := range nums {
			jobs <- n
		}
		close(jobs)
		wg.Wait()
		close(out)
	}()
	s := 0
	for v := range out {
		s += v
	}
	return s
}
func main() { fmt.Println(Pool([]int{1, 2, 3}, 2)) }
