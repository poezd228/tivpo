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
	} else {
		return Search(list[mid:], value)
	}

}

func (list List) BinarySearch(value int) int {
	return Search(*list.listS, value)

}

func main() {
	v := &[]int{1, 3, 5, 6, 8}
	testList := List{listS: v}
	fmt.Print(testList.BinarySearch(5))

}
