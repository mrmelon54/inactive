package inactive

import (
	"testing"
	"time"
)

func TestChannelPassThrough(t *testing.T) {
	z := time.Now()
	a := make(chan int, 10)
	b := make(chan int, 10)
	<-ChannelPassThrough[int](time.Second*2, a, b)
	n := time.Now().Sub(z)
	if n < time.Second*2 {
		t.Fatal("Inactive channel should finish after 2 seconds")
	}
	if n > time.Second*3 {
		t.Fatal("Inactive channel should finish within 3 seconds")
	}
}

func TestChannelPassThroughWithDelay(t *testing.T) {
	z := time.Now()
	a := make(chan int, 10)
	b := make(chan int, 10)
	go func() {
		time.AfterFunc(time.Second/2, func() { a <- 2 })
		time.AfterFunc(time.Second, func() { a <- 1 })
	}()
	<-ChannelPassThrough[int](time.Second*2, a, b)
	n := time.Now().Sub(z)
	if n < time.Second*3 {
		t.Fatal("Inactive channel should finish after 2 seconds")
	}
	if n > time.Second*4 {
		t.Fatal("Inactive channel should finish within 3 seconds")
	}
}
