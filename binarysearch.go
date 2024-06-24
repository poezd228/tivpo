package main

import "fmt"

type List struct {
	listS *[]int
}

func Search(list []int, value int) int {
	mid := len(list) / 2
	if list[mid] == value {
		return mid + 1
	}
	if list[mid] > value {
		return Search(list[:mid], value)
	} else if list[mid] < value {
		return Search(list[mid:], value)
	}
	return mid

}

func BubbleSort(list []int) []int {
	for j := 0; j < len(list); j++ {
		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				list[i], list[i+1] = list[i+1], list[i]
			}

		}
	}

	return list

}
func QuickSort(list []int) []int {
	pivot := len(list) / 2
	left := list[pivot:]
	right := list[:pivot]
	if len(left) > 1 {
		QuickSort(left)
	}
	if len(right) > 1 {
		QuickSort(right)
	}
	return list

}

func (list List) BinarySearch(value int) int {
	return Search(*list.listS, value)

}
func (list List) Sort() {
	fmt.Println(BubbleSort(*list.listS), "Пузыречки")
	fmt.Println(QuickSort(*list.listS), "Быстро")

}

func main() {
	v := &[]int{111, 10, 1, 9, 5, 6, 8, 81}
	testList := List{listS: v}
	fmt.Print(testList.BinarySearch(5), " Поиск ", "\n")
	testList.Sort()

}
