package thread

import (
	"testing"
	"time"
)

func TestCond(t *testing.T) {
	go producer()
	go consumer(1)
	go consumer(2)
	time.Sleep(5)
}