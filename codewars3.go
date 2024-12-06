package main

import (
	"fmt"
	"sync"
)

type Node2 struct {
	value int
	left  *Node2
	right *Node2
}

type Tree2 struct {
	root *Node2
	mu   sync.Mutex // Мьютекс для синхронизации доступа
}

func (t *Tree2) Insert(value int) {
	t.mu.Lock()         // Захват мьютекса
	defer t.mu.Unlock() // Освобождение мьютекса

	newNode := &Node2{value: value}
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

func (t *Tree2) Search(value int) *Node2 {
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
	t := new(Tree2)

	// Запускаем несколько горутин для вставки в дерево
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
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
