package main

import "fmt"

func main() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	quicksort(arr)
	fmt.Println(arr)

}

func quicksort(slice []int) {
	i := 0
	j := len(slice) - 1
	pivot := slice[len(slice)/2]

	if len(slice) <= 1 {
		return
	}
	for {
		for (slice)[i] < pivot {
			i++
		}
		for (slice)[j] > pivot {
			j--
		}
		if i <= j {
			temp := (slice)[i]
			(slice)[i] = (slice)[j]
			(slice)[j] = temp

			i = i + 1
			j = j - 1
		}
		if i > j {
			break
		}

	}
	if j > 0 {
		quicksort(slice[:j+1])
	}
	if len(slice) > i {
		quicksort(slice[i:])
	}
}
