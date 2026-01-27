package main

import "strconv"

func calPoints(operations []string) int {
	stak := []int{}
	res := 0

	for _, opr := range operations {
		switch opr {
		case "+":
			n := len(stak)
			newNum := stak[n-1] + stak[n-2]
			stak = append(stak, newNum)
			res += newNum
		case "D":
			n := len(stak)
			newNum := stak[n-2] * stak[n-1]
			stak = append(stak, newNum)
			res += newNum
		case "C":
			res -= stak[len(stak)-1]
			stak = stak[:len(stak)-1]
		default:
			num, err := strconv.Atoi(opr)
			if err != nil {
				return -1
			}
			stak = append(stak, num)
			res += num
		}
	}
	return res
}
