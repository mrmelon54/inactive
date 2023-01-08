package inactive

import "time"

func ChannelPassThrough[T any](d time.Duration, in, out chan T) chan time.Time {
	t := NewTimer(d)
	go func() {
		for {
			select {
			case z := <-in:
				t.Tick()
				out <- z
			}
		}
	}()
	return t.C
}
