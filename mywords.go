package main

import "fmt"

func SpinWords(str string) [][]int32 {
	list := []rune(str)
	words := make([][]int32, 1)
	var word []int32
	latestI := 0
	for i, v := range list {
		if i == len(list)-1 {
			words = append(words, list[latestI+1:])

		}
		if v == 32 {
			word = list[latestI:i]
			latestI = i
			words = append(words, word)
		}

	}
	return words

}

func main() {
	fmt.Println(SpinWords("абоба хуй"))

}
