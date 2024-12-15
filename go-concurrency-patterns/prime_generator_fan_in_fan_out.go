package main

import (
	"fmt"
	"math/rand/v2"
	"runtime"
	"sync"
	"time"
)


func fanIn[T any](done <- chan int, channels ...<-chan T) <-chan T{
	var wg sync.WaitGroup
	fannedInStream := make(chan T)

	transfer := func(c <- chan T) {
		defer wg.Done()
		for i:= range c {
			select{
			case <- done:
				return
			case fannedInStream <- i:

			}
		}
	}

	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	
	go func(){
		wg.Wait()
		close(fannedInStream)
	}()

	return fannedInStream
}


func primeGeneratorFanInFanOut() {
	start := time.Now()
	done := make(chan int)
	defer close(done)
	randNumFetcher := func() int { return rand.IntN(50000000) }
	randIntStream := repeatFunc(done, randNumFetcher)

	// fan out
	CPUCount := runtime.NumCPU()
	primeFinderChannels := make([]<- chan int, CPUCount)
	for i := 0; i < CPUCount; i++ {
		primeFinderChannels[i] = primeFinder(done, randIntStream)
	}

	// fan in
	primeStream := fanIn(done, primeFinderChannels...)
	
	for rando := range take(done, primeStream, 10) {
		fmt.Println(rando)
	}

	fmt.Println(time.Since(start))
}
