# cwg
Go Channel Wait Group

```
package main 

import "log"

func main() {
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
  log.Printf("results: %v", results)
}
```
