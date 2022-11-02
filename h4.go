package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, m int
	var slice []int
	fmt.Scan(&n)
	slice = make([]int, 0)
	letterMap := make(map[int]int)
	for i := 1; i <= n; i++ {
		fmt.Scan(&m)
		slice = append(slice, m)
	}
	sort.Ints(slice)
	for _, t := range slice {
		if _, ok := letterMap[t]; !ok {
			letterMap[t] = 1
		} else {
			letterMap[t]++
		}
	}
	for key, value := range letterMap {
		fmt.Println(key, value)
	}
}
