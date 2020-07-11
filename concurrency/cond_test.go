package concurrency

import (
	"testing"
	"time"
)

func TestCond(t *testing.T) {
	go condProducer()
	go condConsumer(1)
	go condConsumer(2)
	time.Sleep(5)
}
