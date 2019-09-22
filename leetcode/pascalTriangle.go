package leetcode

func Generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i:=0;i<numRows ;i++  {
		for j := 0; j < numRows; j++ {
			triangle[i] = append(triangle[i], 0)
		}
	}
	for i := 0; i < numRows; i++ {
		triangle[i][0] = 1
		triangle[i][i] = 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1]+triangle[i-1][j]
		}
		triangle[i] = triangle[i][:i+1]
	}
	return triangle
}
