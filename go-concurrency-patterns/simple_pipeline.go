package main

import (
	"fmt"
)

func someFunc(done <-chan bool) {
	for {
		select {
		case <-done:
			return
		default:
			fmt.Println("DOING WORK")
		}
	}
}

func SliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, item := range nums {
			out <- item
		}
		close(out)
	}()
	return out
}

func sq(dataChannel <-chan int) <-chan int {

	sq := make(chan int)
	go func() {
		for item := range dataChannel {
			sq <- item * item
		}
		close(sq)
	}()
	return sq
}

func simplePipeline() {

	// input
	nums := []int{2, 3, 4, 7, 1}

	// stage1
	dataChannel := SliceToChannel(nums)

	// stage2
	finalChannel := sq(dataChannel)
	// stage 3
	for n := range finalChannel {
		println(n)
	}
}
