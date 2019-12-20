package ui

import "time"

// InputChan represent a throttled channel of input reading each 8ms.
type InputChan struct {
	ci     chan Input
	ticker <-chan time.Time
}

// NewInputChan returns a valid InputChan.
func NewInputChan() InputChan {
	return InputChan{
		ci:     make(chan Input),
		ticker: time.Tick(8 * time.Millisecond),
	}
}

func (c InputChan) Read(f func(Input)) {
	var lastInput Input
	for {
		select {
		case inp := <-c.ci:
			lastInput = inp
		case <-c.ticker:
			f(lastInput)
		}
	}
}
