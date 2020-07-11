package concurrency

import (
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	go semProducer()
	go semConsumer(1)
	go semConsumer(2)
	time.Sleep(5)
}
