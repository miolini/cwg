package cwg

import (
	"sync/atomic"
	"time"
)

type CWG struct {
	channel chan interface{}
	counter int64
}

func New() *CWG {
	cwg := new(CWG)
	cwg.channel = make(chan interface{})
	return cwg
}

func (cwg *CWG) Add(a int64) int64 {
	return atomic.AddInt64(&cwg.counter, a)
}

func (cwg *CWG) Counter() int64 {
	return atomic.LoadInt64(&cwg.counter)
}

func (cwg *CWG) Done(v interface{}) {
	cwg.channel <- v
}

func (cwg *CWG) WaitWithoutResults() {
	stop := false
	for !stop {
		select {
		case <- cwg.channel:
			counter := cwg.Add(-1)
			if counter == 0 {
				stop = true
				break
			}
		}
	}
}

func (cwg *CWG) Wait() (msgs []interface{}) {
	msgs = []interface{}{}
	stop := false
	for !stop {
		select {
		case msg := <- cwg.channel:
			msgs = append(msgs, msg)
			counter := cwg.Add(-1)
			if counter == 0 {
				stop = true
				break
			}
		}
	}
	return
}

func (cwg *CWG) WaitWithTimeout(timeout time.Duration) (ok bool, msgs []interface{}) {
	msgs = []interface{}{}
	for {
		select {
		case msg := <- cwg.channel:
			msgs = append(msgs, msg)
			counter := cwg.Add(-1)
			if counter == 0 {
				ok = true
				break
			}
		case <- time.After(timeout):
			break
		}
	}
	return
}