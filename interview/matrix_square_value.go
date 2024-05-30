package interview

// 题目：计算方阵的主对角线和次对角线的绝对值的差值
// 思路：根据方阵的宽度进行for循环通过 array[i][i] 获取每个对角线的值然后相加，
//
//	次对角线则通过array[i][n-i-1]获取，然后再相减得出结果

func matrixSquareXvalue(matrix [][]int) int {
	n := len(matrix)
	primaryDiagonalSum := 0
	secondaryDiagonalSum := 0
	for i := 0; i < n; i++ {
		primaryDiagonalSum += matrix[i][i]
		secondaryDiagonalSum += matrix[i][n-i-1]
	}

	result := primaryDiagonalSum - secondaryDiagonalSum
	if result < 0 {
		return -result
	}
	return result
}
