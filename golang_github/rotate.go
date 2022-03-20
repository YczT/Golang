package main

import "fmt"

/*
给你一幅由 N × N 矩阵表示的图像，其中每个像素的大小为 4 字节。请你设计一种算法，将图像旋转 90 度。

不占用额外内存空间能否做到？
*/
func rotate(matrix [][]int) [][]int {
	length := len(matrix)
	if length == 0 {
		return nil
	}
	//行对称变换
	for i := 0; i < length/2; i++ {
		matrix[i], matrix[length-1-i] = matrix[length-1-i], matrix[i]
	}
	//主对角线对称变换
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	return matrix
}
func main() {
	var test_e = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	//ans1 := test_e[0]
	//ans2 := test_e[3]
	//fmt.Printf("%v\n", ans1)
	//fmt.Printf("%v\n", ans2)
	//test_e[0], test_e[3] = test_e[3], test_e[0]

	var ans = rotate(test_e)
	//ans1 = test_e[0]
	//ans2 = test_e[3]
	//fmt.Printf("%v\n", ans1)
	fmt.Printf("%v\n", ans)
}
