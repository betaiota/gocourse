package tasks

import "fmt"

func intersection(arr1, arr2 []string) []string {
	bucket := make(map[string]bool)
	for _, i := range arr1 {
		bucket[i] = true
	}

	var commonElements []string
	added := make(map[string]bool)

	for _, j := range arr2 {
		if bucket[j] && !added[j] {
			commonElements = append(commonElements, j)
			added[j] = true
		}
	}

	return commonElements
}

func LabEighteen() {

	//Даны читатели книг
	var readers_books = [...]string{"id3", "id5", "id9", "id8", "id2", "id1"}

	//Даны читатели газет
	var readers_magazines = [...]string{"id8", "id2", "id1", "id4", "id6", "id7", "id10"}

	both := intersection(readers_books[:], readers_magazines[:])
	fmt.Println("Those who read both magazines and books:", both)
}
