package main

import (
	"fmt"
	"sync"
)


func confinement() {

	var wg sync.WaitGroup

	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	result := make([]int, len(input))

	for i, data := range input {
		wg.Add(1)
		go processData(&wg, &result[i], data)
	}

	wg.Wait()

	fmt.Println(result)
}

func processData(wg *sync.WaitGroup, result *int, data int) {
	defer wg.Done()
	processedData := data * 2
	*result = processedData
}


// Mutex way
//var lock sync.Mutex
// func processData(wg *sync.WaitGroup, result *[]int, data int) {
// 	defer wg.Done()
// 	processedData := data * 2
// 	lock.Lock()
// 	*result = append(*result, processedData)
// 	lock.Unlock()
// }