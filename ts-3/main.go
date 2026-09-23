// Уявіть, що є семафор ємністю 1 горутина.
// Ми намагаємося запустити три фонові задачі паралельно.
// Що саме виведе програма в консоль і скільки завдань буде реально виконано?

package main

import (
	"fmt"
	"time"
)

func main() {
	sem := make(chan struct{}, 1)

	for i := 1; i <= 3; i++ {
		go func(id int) {
			select {
			case sem <- struct{}{}:
				fmt.Printf("Worker %d started\n", id)
				time.Sleep(100 * time.Millisecond)
				<-sem
				fmt.Printf("Worker %d finished\n", id)
			default:
				fmt.Printf("Worker %d skipped\n", id)
			}
		}(i)
	}

	time.Sleep(1 * time.Second)
}
