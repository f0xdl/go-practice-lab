package tasks

import (
	"sort"
	"sync"
)

/*
	ProcessData

# TASK 2

🧠 Задача: Фан-аут и сбор результатов
Напиши функцию ProcessData, которая:

Принимает на вход слайс []int

# Создаёт 3 горутины, каждая из которых читает числа из входного канала и возводит их в квадрат

# Все результаты собираются в выходной канал

# Функция возвращает отсортированный слайс квадратов чисел

Уточнения:

Используй fan-out (делим работу между несколькими горутинами)

# Обязательно использовать каналы и WaitGroup

Не забудь закрыть каналы правильно, чтобы избежать утечек
*/
func ProcessData(workers int, val []int) []int {
	chOut := make(chan int, len(val))

	chIn := make(chan int, len(val))
	var result []int

	wg := sync.WaitGroup{}
	wg.Add(workers)
	for _, v := range val {
		chIn <- v
	}
	close(chIn)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			for {
				res, ok := <-chIn
				if !ok {
					return
				}
				res *= res
				chOut <- res
			}
		}()
	}
	wg.Wait()
	close(chOut)

	//for {
	//	if res, ok := <-chOut; ok {
	//		result = append(result, res)
	//	} else {
	//		break
	//	}
	//}
	for res := range chOut {
		result = append(result, res)
	}
	//return SortingSlice(result)
	sort.Ints(result)
	return result
}

func SortingSlice(result []int) []int {
	for i := 0; i < len(result); i++ {
		for j := i + 1; j < len(result); j++ {
			if result[i] > result[j] {
				result[i], result[j] = result[j], result[i]
			}
		}
	}
	return result
}
