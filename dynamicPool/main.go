package main

import (
	"context"
	"fmt"
	"go-initial/dynamicPool/dwp"
	"time"
)

func processData(workerId int, data string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Worker %d processing %s\n", workerId, data)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Min: 2 воркери, Max: 6 воркерів, IdleTimeout: 1 секунда, Черга: 2 задачі
	pool := dwp.New(2, 6, 1*time.Second, processData)
	fmt.Printf("--- Початковий стан: %d активних воркерів ---\n", pool.ActiveWorkersCount())

	// 1. Подаємо сплеск навантаження (10 задач поспіль)
	fmt.Println("\n>>> Відправляємо 10 задач (СПЛЕСК НАВАНТАЖЕННЯ)...")

	for i := 1; i <= 10; i++ {
		pool.Submit(ctx, fmt.Sprintf("Job #%d", i))
	}

	fmt.Printf("--- Активних воркерів під час піку: %d ---\n", pool.ActiveWorkersCount())

	// 2. Чекаємо, поки задачі обробляться і настане період спокою
	fmt.Println("\n>>> Чекаємо 3 секунди (Період бездіяльності)...")
	time.Sleep(3 * time.Second)
	fmt.Printf("--- Після простою активних воркерів: %d ---\n", pool.ActiveWorkersCount())

	// Зупиняємо пул
	pool.Stop()
	fmt.Println("\n>>> Пул зупинено.")

}
