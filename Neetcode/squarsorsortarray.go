package main

func Squares(arr []int) []int {

	arrayLen := len(arr)
	l, r := 0, arrayLen-1
	position := arrayLen - 1

	newArray := make([]int, arrayLen)

	for l <= r {

		leftSquar := arr[l] * arr[l]
		rightSquar := arr[r] * arr[r]

		if leftSquar > rightSquar {
			newArray[position] = leftSquar
			l++
		} else {
			newArray[position] = rightSquar
			r--
		}
		position--
	}
	return newArray
}
