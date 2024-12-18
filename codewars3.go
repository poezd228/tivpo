package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

type Node3 struct {
	value int
	left  *Node3
	right *Node3
}

type Tree3 struct {
	root *Node3
	mu   sync.Mutex // Мьютекс для синхронизации доступа
}

func (t *Tree3) Insert(value int) {
	t.mu.Lock()         // Захват мьютекса
	defer t.mu.Unlock() // Освобождение мьютекса

	newNode := &Node3{value: value}
	if t.root == nil {
		t.root = newNode
		return
	}
	current := t.root
	for {
		if value < current.value {
			if current.left == nil {
				current.left = newNode
				return
			}
			current = current.left
		} else {
			if current.right == nil {
				current.right = newNode
				return
			}
			current = current.right
		}
	}
}

func (t *Tree3) Search(value int) *Node3 {
	t.mu.Lock()
	defer t.mu.Unlock()

	if t.root == nil {
		return nil
	}
	current := t.root
	for {
		if value == current.value {
			return current
		} else if value < current.value {
			current = current.left
		} else {
			current = current.right
		}
		if current == nil {
			return nil
		}
	}
}

func main() {
	t := new(Tree3)
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	slice := make([]byte, 1<<30)
	fmt.Println("Memory allocated:", len(slice))
	// Запускаем несколько горутин для вставки в дерево
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		// 1 GB

		// Наполнение среза для "активного" использования памяти

		slice[i] = 255

		// Пауза, чтобы наблюдать за увеличением потребления памяти
		time.Sleep(10 * time.Second)
		time.Sleep(10000000000)
		wg.Add(1)

		go func(val int) {
			defer wg.Done()
			t.Insert(val)
		}(i)
	}

	wg.Wait()

	// Теперь мы можем безопасно искать элементы в дереве
	node := t.Search(3)
	fmt.Println(node) // Проверка на вывод результатов

	// Включите динамический анализ для обнаружения гонок данных
	// Для этого используйте команду go run -race
}
