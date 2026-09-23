// Уявіть, що воркер обробляє вхідні повідомлення з каналу jobs всередині циклу for-select.
// Нам потрібно реалізувати логіку: якщо протягом 500 мілісекунд у канал jobs не надходить
// жодної нової задачі, або якщо приходить сигнал скасування з каналу done, воркер має миттєво припинити
// роботу та завершити горутину.
// Який примітив із пакету time найкраще використати всередині select для реалізації цього таймауту і як виглядатиме структура цього select?

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	numbs := []int{1, 2, 3, 4, 5}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	chanNumbs := generate(ctx, numbs...)

	worker(ctx, chanNumbs)

}

func generate(ctx context.Context, numbs ...int) <-chan int {

	out := make(chan int)

	go func() {
		defer close(out)

		for _, i := range numbs {
			select {
			case <-ctx.Done():
				return
			case out <- i:
			}
		}
	}()

	return out

}

func worker(ctx context.Context, jobs <-chan int) {
	ticker := time.NewTimer(2 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			if !ticker.Stop() {
				select {
				case <-ticker.C:
				default:
				}
			}
			ticker.Reset(500 * time.Millisecond)
			fmt.Println("Worker started processing a job: ", job)
		}

	}
}
