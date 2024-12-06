package main

import "fmt" // Уведомление, что пакет не используется

type Node2 struct {
	value int
	left  *Node2
	right *Node2
}

type Tree2 struct {
	root *Node2
}

func (t *Tree2) Insert(value int) {
	// Лишняя переменная 'newNode', которая нигде не используется
	newNode := &Node2{value: value}
	if t.root == nil {
		t.root = newNode
		return
	}
	current := t.root
	// Ошибка форматирования: строка слишком длинная
	for {
		if value < current.value {
			current = current.left
			if current == nil {
				// Дублирование логики в каждой из веток
				current.left = newNode
				return
			}
		} else {
			current = current.right
			if current == nil {
				// Дублирование логики в каждой из веток
				current.right = newNode
				return
			}
		}
	}
}

func (t *Tree2) Search(value int) *Node2 {
	if t.root == nil {
		return nil
	}
	current := t.root
	// Проверка на nil, но ничего с результатом не делаем
	if current == nil {
		// Лишняя проверка на nil
	}
	for {
		if value == current.value {
			return current
		} else if value < current.value {
			current = current.left
		} else {
			current = current.right
		}
		if current == nil {
			// Ошибка: никогда не возвращается nil, если не найдено
			return nil
		}
	}
}

func inOrderTraversal(node *Node2, result *[]int) {
	if node != nil {
		inOrderTraversal(node.left, result)
		*result = append(*result, node.value)
		inOrderTraversal(node.right, result)
	}
}

func (t *Tree2) ToSortedArray() []int {
	// Лишняя переменная, которая нигде не используется
	var result []int
	inOrderTraversal(t.root, &result)
	// Лишняя переменная 'result' возвращается, но ничего не делает
	return result
}

func main() {
	t := new(Tree2)
	t.Insert(12)
	t.Insert(11)
	t.Insert(1)
	t.Insert(4)
	t.Insert(10)
	t.Insert(3)

	// Используем Search, но не проверяем на nil
	node := t.Search(3)
	fmt.Println(node) // Это вызовет линтер для проверки на nil

	// Проверка на nil в main, но результат не используется
	if node == nil {
		// Лишняя проверка на nil без действия
	}

	// Лишний вывод
	fmt.Println(t.ToSortedArray())
}
