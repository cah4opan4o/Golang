package main

import (
	"fmt"
	"sync"
	"time"
)

// Структура турникета
type turniketCounter struct {
	count int //количество срабатываний
	mutex sync.Mutex
}

// функция для увелечения счётчика
func (t *turniketCounter) Growth() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	t.count++ // увелечение счётчика
}

// функция для получения информации о количестве человек
func (t *turniketCounter) Value() int {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	return t.count
}

func main() {
	// Создаем экземпляр счетчика
	counter := turniketCounter{}

	// Функция для прохода через турникет
	passThrough := func() {
		counter.Growth()
		fmt.Printf("Человек прошел через турникет. Текущее количество прошедших: %d\n", counter.Value())
		time.Sleep(time.Second) // Имитация времени прохода
	}

	// запуск нескольких горутин
	for i := 1; i <= 10; i++ {
		go passThrough()
	}

	time.Sleep(10 * time.Second)
	fmt.Printf("Всего человек прошло через турникет: %d\n", counter.Value())
}
