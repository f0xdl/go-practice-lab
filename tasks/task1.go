package tasks

import "sync"

/*
	CounterCheck

Напиши функцию на Go, которая запускает N горутин.
Каждая горутина увеличивает общий счётчик на 1, 1000 раз.
Твоя задача — убедиться, что по завершению всех горутин значение счётчика равно N * 1000.

Условия:
# Использовать sync.WaitGroup для ожидания завершения
# Использовать sync.Mutex для защиты счётчика
*/
func CounterCheck(workers int) bool {
	counter := 0
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()

			for i := 0; i < 1000; i++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	//fmt.Println("counter=", counter)
	if workers*1000 == counter {
		return true
	} else {
		return false
	}
}
