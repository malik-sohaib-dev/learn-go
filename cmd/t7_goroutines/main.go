package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var m = sync.Mutex{} // Mutex: Mutual Exclusion Lock - It is used to prevent multiple goroutines from accessing the same resource at the same time
var wg = sync.WaitGroup{}

var sampleDB = []string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth"}
var result = make([]string, 0, len(sampleDB))

func main() {
	// Concurrency using goroutines: Goroutines are lightweight threads that are managed by Go runtime
	// Goroutines are functions that run concurrently with other functions

	// Sequential execution: If we run the below code, it will run sequentially
	// t0 := time.Now()
	// for i := 0; i < len(sampleDB); i++ {
	// 	dbCallSequential(i)
	// }
	// fmt.Println("Sequenctial Time taken: ", time.Since(t0))

	fmt.Println("-----------------------------")
	// Concurrency using goroutines: If we add go keyword before the function call, it will run concurrently
	// Although, this will just spawn the goroutines in the background and not wait for them to finish executing
	// So we need a way to wait for all goroutines to finish executing
	// We can use WaitGroup for this
	// WaitGroup is a struct that has 3 methods: Add, Done, Wait
	// To use WaitGroup, we need to import sync package
	// WaitGroup is just a counter that keeps track of how many goroutines are running
	// Steps: 1. Create a WaitGroup 2. Increment the counter when a goroutine is spawned 3. Decrement the counter when the goroutine is done 4. Wait for all goroutines to finish
	// This implementation is not efficient as if multiple routines try to append to the same slice, it could cause data corruption and many more issues
	// So we need to use Mutex to prevent multiple goroutines from accessing the same resource at the same time
	// Mutex is a lock that is used to prevent multiple goroutines from accessing the same resource at the same time
	// Mutex is a struct that has 2 methods: Lock, Unlock
	// To use Mutex, we need to import sync package
	// Steps: 1. Create a Mutex 2. Lock the resource before accessing it 3. Unlock the resource after accessing it
	t1 := time.Now()

	for i := 0; i < len(sampleDB); i++ {
		wg.Add(1) // Increment the counter
		go dbCall(i)
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Concurrent Time taken: ", time.Since(t1))
	fmt.Println("Result: ", result)
}

func dbCall(i int) {
	var randomTime = 2000
	time.Sleep(time.Duration(randomTime) * time.Millisecond)
	m.Lock() // Lock the resource
	result = append(result, sampleDB[i])
	m.Unlock() // Unlock the resource
	fmt.Printf("Iteration: %d: - %v\n", i, sampleDB[i])
	wg.Done() // Decrement the counter
}

func dbCallSequential(i int) {
	var randomTime = rand.Float32() * 2000
	time.Sleep(time.Duration(randomTime) * time.Millisecond)
	fmt.Printf("Iteration: %d: - %v\n", i, sampleDB[i])
}
