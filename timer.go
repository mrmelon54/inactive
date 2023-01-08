package inactive

import "time"

type Timer struct {
	C chan time.Time
	t *time.Timer
	d time.Duration
	r chan bool
}

func NewTimer(d time.Duration) *Timer {
	i := &Timer{
		C: make(chan time.Time, 1),
		t: time.NewTimer(d),
		d: d,
		r: make(chan bool, 1),
	}
	go func() {
	keepWaiting:
		for {
			select {
			case n := <-i.r:
				if !i.t.Stop() {
					<-i.t.C
				}
				if n {
					break keepWaiting
				}
				i.t.Reset(i.d)
			case n := <-i.t.C:
				i.C <- n
			}
		}
	}()
	return i
}

func (i *Timer) Tick() {
	i.r <- false
}

func (i *Timer) Stop() {
	i.r <- true
}
