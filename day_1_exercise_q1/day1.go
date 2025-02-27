package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	Rows, Cols int
	Elements   [][]int
}

func (m *Matrix) getRows() int {
	return m.Rows
}

func (m *Matrix) getCols() int {
	return m.Cols
}

func (m *Matrix) setElement(i, j, val int) {
	m.Elements[i][j] = val
	for _, row := range m.Elements {
		fmt.Println(row)
	}
}

func (m *Matrix) addMatrix(other Matrix) Matrix {
	if m.Rows != other.Rows || m.Cols != other.Cols {
		fmt.Println("Matrices dimensions do not match")
		return Matrix{}
	}
	result := Matrix{
		Rows:     m.Rows,
		Cols:     m.Cols,
		Elements: make([][]int, m.Rows),
	}

	for i := 0; i < m.Rows; i++ {
		result.Elements[i] = make([]int, m.Cols)
		for j := 0; j < m.Cols; j++ {
			result.Elements[i][j] = m.Elements[i][j] + other.Elements[i][j]
		}
	}
	return result
}

func main() {
	s := Matrix{
		Rows: 3,
		Cols: 3,
		Elements: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	}
	s1 := Matrix{
		Rows: 3,
		Cols: 3,
		Elements: [][]int{
			{9, 8, 7},
			{6, 5, 4},
			{3, 2, 1},
		},
	}
	fmt.Println("Rows:", s.getRows())
	fmt.Println("Columns:", s.getCols())
	s.setElement(2, 2, 10)
	result := s.addMatrix(s1)

	fmt.Println("Resultant Matrix after addition:")
	for _, row := range result.Elements {
		fmt.Println(row)
	}

	jsonData, err := json.Marshal(result)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	fmt.Println("Matrix as JSON:", string(jsonData))
}
