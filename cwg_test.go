package cwg

import (
	"testing"
	"time"
	"math/rand"
)

func TestChannelWaitGroup(t *testing.T) {
	cwg := New()
	var num int64 = 10000
	cwg.Add(num)
	var i int64
	for i=0; i < cwg.Counter(); i++ {
		go func(num int64, cwg *CWG) {
			defer cwg.Done(num)
			time.Sleep(time.Millisecond * time.Duration(rand.Int() % 5000))
		}(i, cwg)
	}
	results := cwg.Wait()
	if int64(len(results)) != num {
		t.Fatalf("results len not equals: %d != %d", num, len(results))
	}
}