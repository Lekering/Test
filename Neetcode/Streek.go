package main

import "fmt"

type Repository interface {
	GetData(id int) (string, error)
	SaveData(id int, data string) error
}

func longestConsecutive(nums ...int) int {
	res := 0
	if len(nums) == 0 {
		return res
	}

	NumsMap := make(map[int]struct{})

	for _, num := range nums {
		NumsMap[num] = struct{}{}
	}

	for _, num := range nums {
		streek, count := 0, num

		for {
			_, ok := NumsMap[count]
			if !ok {
				break
			}
			streek++
			count++
		}
		if res < streek {
			res = streek
		}
	}
	fmt.Println(res)
	return res
}
