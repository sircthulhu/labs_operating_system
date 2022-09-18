package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// printMatrix печатает матрицу на экран
func printMatrix(m *[][]int) {
	rows := len(*m)
	cols := len((*m)[0])
	
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			fmt.Print((*m)[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

// createRandomMatrix создает матрицу случайных чисел размером n строк, m столбцов
func createRandomMatrix(n, m int) *[][]int {
	matrix := make([][]int, n)
	
	// Заполняем параллельно матрицу случайным образом числами [0;100)
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		col := i
		go func(col *[]int) {
			defer wg.Done()
			
			*col = make([]int, m)
			for j := 0; j < m; j++ {
				(*col)[j] = rand.Intn(100)
			}
		}(&matrix[col])
	}
	wg.Wait()
	
	return &matrix
}

func calculateMinElements(matrix *[][]int) *[]int {
	rows := len(*matrix)
	cols := len((*matrix)[0])
	
	mins := make([]int, cols)
	wg := sync.WaitGroup{}
	for i := 0; i < cols; i++ {
		wg.Add(1)
		col := i
		go func(col int, result *int) {
			defer wg.Done()
			
			*result = (*matrix)[0][col]
			for j := 0; j < rows; j++ {
				if (*matrix)[j][col] < *result {
					*result = (*matrix)[j][col]
				}
			}
		}(col, &mins[i])
	}
	wg.Wait()
	
	return &mins
}

func main() {
	n := 10
	m := 15
	
	// matrix это двумерный массив целых чисел
	matrix := createRandomMatrix(n, m)
	
	fmt.Println("Созданная матрица:")
	printMatrix(matrix)
	
	// Считаем минимальные значения столбцов
	fmt.Println("Вычисляем минимумы по столбцам...")
	mins := calculateMinElements(matrix)
	fmt.Printf("Минимальные элементы по столбцам: %v\n", *mins)
	
	min := (*mins)[0]
	minIdx := 0
	for i := 1; i < m; i++ {
		if min > (*mins)[i] {
			min = (*mins)[i]
			minIdx = i
		}
	}
	fmt.Printf("Минимальный элемент - %d. Столбец этого элемента - %d\n", min, minIdx)
	
}
