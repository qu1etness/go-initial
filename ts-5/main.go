// Завдання 2: Проектування триетапного конвеєра
// Вам необхідно реалізувати конвеєр обробки послідовності чисел:
// Етап 1 (generate): Приймає список цілих чисел ...int, записує їх у канал і закриває його
// .
// Етап 2 (filterOdds): Приймає вхідний канал, відфільтровує парні числа та пропускає далі в аутпут-канал лише непарні
// .
// Етап 3 (multiplyByTen): Приймає відфільтровані числа, множить кожне на 10 та передає далі.
// Опишіть (або напишіть сигнатури та тіла функцій), як закриття каналу на етапі generate призведе до каскадного та безпечного завершення всього конвеєра
package main

import (
	"context"
	"fmt"
)

func main() {

	numbers := []int{1, 2, 3, 4, 5}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	chanNumbs := generate(ctx, numbers...)
	chanOdds := filterOdds(ctx, chanNumbs)
	chanResult := multiplyByTen(ctx, chanOdds)

	for v := range chanResult {
		fmt.Println(v)
	}

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

func filterOdds(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				if v%2 != 0 {
					select {
					case <-ctx.Done():
						return
					case out <- v:
					}
				}
			}
		}

	}()

	return out
}

func multiplyByTen(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				select {
				case <-ctx.Done():
					return
				case out <- v * 10:
				}
			}
		}
	}()

	return out
}
