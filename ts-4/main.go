// Перед вами функція-генератор gen, яка асинхронно продукує нескінченну послідовність
// чисел. Головна горутина main зчитує лише перші 5 чисел і завершує обробку.
// Знайдіть помилку витоку горутини та перепишіть код, додавши сигнальний канал done,
// щоб генератор коректно завершував роботу після того, як main припинить читання.

package main

import "fmt"

func gen(done chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; ; i++ {
			select {
			case out <- i: // Де виникає блокування горутини?
			case <-done:
				return
			}
		}
	}()
	return out
}

func main() {
	done := make(chan struct{})
	defer close(done)
	ch := gen(done)
	for i := 0; i < 9; i++ {
		fmt.Println(<-ch)
	}
	// Що відбувається з горутиною всередині gen() після завершення цього циклу?
}
