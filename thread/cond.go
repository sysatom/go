package thread

import (
	"fmt"
	"sync"
)

const MAX = 12

var buffer [MAX]int
var fillPtr = 0
var usePtr = 0
var count = 0

func put(value int) {
	buffer[fillPtr] = value
	fillPtr = (fillPtr + 1) % MAX
	count++
}

func get() int {
	tmp := buffer[usePtr]
	usePtr = (usePtr + 1) % MAX
	count--
	return tmp
}

const ProducerLoops = 10
const ConsumerLoops = 5

var mutex = sync.Mutex{}
var emptyCond = sync.NewCond(&mutex)
var fillCond = sync.NewCond(&mutex)

func producer() {
	for i := 0; i < ProducerLoops; i++ {
		mutex.Lock()
		for count == MAX {
			emptyCond.Wait()
		}
		put(i)
		fillCond.Signal()
		mutex.Unlock()
	}
}

func consumer(n int) {
	for i := 0; i < ConsumerLoops; i++ {
		mutex.Lock()
		for count == 0 {
			fillCond.Wait()
		}
		tmp := get()
		emptyCond.Signal()
		mutex.Unlock()
		fmt.Println(n, tmp)
	}
}

