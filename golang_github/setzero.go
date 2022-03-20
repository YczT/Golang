package main

import "fmt"

/*
编写一种算法，若M × N矩阵中某个元素为0，则将其所在的行与列清零。
*/
func setZeroes(matrix [][]int) [][]int {
	length_x := len(matrix)
	if length_x < 0 {
		return matrix
	}
	length_y := len(matrix[0])
	xmap := make(map[int]bool)
	ymap := make(map[int]bool)
	for i := 0; i < length_x; i++ {
		for j := 0; j < length_y; j++ {
			if matrix[i][j] == 0 {
				xmap[i], ymap[j] = true, true
			}
		}
	}
	for i := 0; i < length_x; i++ {
		for j := 0; j < length_y; j++ {
			if xmap[i] || ymap[j] {
				matrix[i][j] = 0
			}
		}
	}

	return matrix
}

func main() {
	var test_e = [][]int{
		{1, 1, 1, 0},
		{1, 0, 1, 1},
		{1, 1, 1, 1},
	}

	//ans := len(test_e[0])
	ans := setZeroes(test_e)
	fmt.Println(ans)
}
