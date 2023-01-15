package inactive

import (
	"testing"
	"time"
)

func TestInactiveTimer(t *testing.T) {
	z := time.Now()
	a := NewTimer(time.Second * 2)
	<-a.C
	n := time.Now().Sub(z)
	if n < time.Second*2 {
		t.Fatal("Inactive timer should finish after 2 seconds")
	}
	if n > time.Second*3 {
		t.Fatal("Inactive timer should finish within 3 seconds")
	}
}

func TestInactiveTimerWithDelay(t *testing.T) {
	z := time.Now()
	a := NewTimer(time.Second * 2)
	go func() {
		time.AfterFunc(time.Second*1, func() { a.Tick() })
	}()
	<-a.C
	n := time.Now().Sub(z)
	if n < time.Second*3 {
		t.Fatal("Inactive timer should finish after 3 seconds")
	}
	if n > time.Second*4 {
		t.Fatal("Inactive timer should finish within 4 seconds")
	}
}

func TestInactiveTimerWithLongerDelay(t *testing.T) {
	z := time.Now()
	a := NewTimer(time.Second * 2)
	go func() {
		time.AfterFunc(time.Second*3, func() { a.Tick() })
	}()
	<-a.C
	<-a.C
	n := time.Now().Sub(z)
	if n < time.Second*5 {
		t.Fatal("Inactive timer should finish after 5 seconds")
	}
	if n > time.Second*6 {
		t.Fatal("Inactive timer should finish within 6 seconds")
	}
}
