package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func contexMain() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	generator := func(dataItem string, stream chan interface{}) {
		for {
			select {
			case <-ctx.Done():
				return
			case stream <- dataItem:
			}
		}
	}

	infiniteApples := make(chan interface{})
	go generator("apple", infiniteApples)

	infiniteOranges := make(chan interface{})
	go generator("orange", infiniteOranges)

	infinitePineapples := make(chan interface{})
	go generator("pineapple", infinitePineapples)

	wg.Add(1)
	go funcApple(ctx, &wg, infiniteApples)

	funcOrange := genericFunc
	funcPinneaple := genericFunc

	wg.Add(1)
	go funcOrange(ctx, &wg, infiniteOranges)
	wg.Add(1)
	go funcPinneaple(ctx, &wg, infinitePineapples)

	wg.Wait()

}

func funcApple(ctx context.Context, parentWg *sync.WaitGroup, stream <-chan interface{}) {

	defer parentWg.Done()
	var wg sync.WaitGroup

	doWork := func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case data, ok := <-stream:
				if !ok {
					fmt.Println("channel closed")
					return
				}
				fmt.Println(data)
			}
		}
	}

	newCtx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go doWork(newCtx)
	}

	wg.Wait()
}

func genericFunc(ctx context.Context, wg *sync.WaitGroup, stream <-chan interface{}) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case data, ok := <-stream:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println(data)
		}
	}
}
