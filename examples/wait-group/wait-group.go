package waitgroup

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second * 2)
	fmt.Printf("Worker %d done\n", id)
}

// wait group
func WaitGroup() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {

		wg.Add(1)
		i := i

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()

}
