package main

import (
	"fmt"
	"sync"
)

func orDone(done <-chan struct{}, dataCh <-chan interface{}) <-chan interface{} {
	relayStream := make(chan interface{})
	go func() {
		defer close(relayStream)
		for {
			select {
			case <-done:
				// Exit if the done channel is closed
				return
			case v, ok := <-dataCh:
				if !ok {
					// Exit if the data channel is closed
					return
				}
				select {
				case relayStream <- v: // Send the value to result channel
				case <-done: // Exit if done signal arrives while sending
				}
			}
		}
	}()
	return relayStream
}

func consumeCows(done <-chan struct{}, cows <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for cow := range orDone(done, cows) {
		// do some complex logic
		fmt.Println(cow)
	}
}

func consumeLamb(done <-chan struct{}, lambs <-chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for lamb := range orDone(done, lambs) {
		// do some complex logic
		fmt.Println(lamb)
	}
}

func ordoneComsumers() {

	var wg sync.WaitGroup
	done := make(chan struct{})
	defer close(done)

	cows := make(chan interface{}, 100)
	lambs := make(chan interface{}, 100)

	go func() {
		for {
			select {
			case <-done:
				return
			case cows <- "moo":
			}
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				return
			case lambs <- "meuh":
			}
		}
	}()

	wg.Add(1)
	go consumeCows(done, cows, &wg)
	wg.Add(1)
	go consumeLamb(done, lambs, &wg)

	wg.Wait()

}
