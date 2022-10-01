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

// Процедура calculateSort выполняет сортировку подсчетом переданного массива
func calculateSort(arr *[]int) {
	// maxNumberInArray это максимальное число в массиве arr
	var maxNumberInArray int = 99

	c := make([]int, maxNumberInArray+1) // Числа в исходном массиве [0;99]
	length := len(*arr)
	for i := 0; i < length; i++ {
		c[(*arr)[i]]++
	}

	idx := 0
	for i := 0; i < maxNumberInArray+1; i++ {
		for j := 0; j < c[i]; j++ {
			(*arr)[idx] = i
			idx++
		}
	}
}

func main() {
	n := 3
	m := 3

	// matrix это двумерный массив целых чисел
	matrix := createRandomMatrix(n, m)

	fmt.Println("Созданная матрица:")
	printMatrix(matrix)

	fmt.Println("Сортируем матрицу по строкам...")
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()
			calculateSort(&((*matrix)[row]))
		}(i)
	}

	fmt.Println("Отсортированная по строкам матрица:")
	printMatrix(matrix)
}

/*
Тестовый запуск:
Созданная матрица:
81 87 47 59 81 18 25 40 56 0 94 11 62 89 28
8 87 31 29 56 37 31 85 26 13 90 94 63 33 76
87 49 28 18 84 3 24 47 12 32 16 39 40 86 51
94 26 2 81 79 66 70 93 86 19 81 52 75 85 10
11 45 37 6 95 66 28 58 47 47 87 88 90 15 41
55 51 10 5 56 66 28 61 2 83 46 63 76 2 18
47 94 77 63 96 20 23 53 37 33 41 59 33 43 91
2 78 36 46 7 40 3 52 43 5 98 25 51 15 57
87 10 10 85 90 32 98 53 91 82 84 97 67 37 71
74 47 78 24 59 53 57 21 89 99 0 5 88 38 3
Сортируем матрицу по строкам...
Отсортированная по строкам матрица:
0 11 18 25 28 40 47 56 59 62 81 81 87 89 94
8 13 26 29 31 31 33 37 56 63 76 85 87 90 94
3 12 16 18 24 28 32 39 40 47 49 51 84 86 87
2 10 19 26 52 66 70 75 79 81 81 85 86 93 94
6 11 15 28 37 41 45 47 47 58 66 87 88 90 95
2 2 5 10 18 28 46 51 55 56 61 63 66 76 83
20 23 33 33 37 41 43 47 53 59 63 77 91 94 96
2 3 5 7 15 25 36 40 43 46 51 52 57 78 98
10 10 32 37 53 67 71 82 84 85 87 90 91 97 98
0 3 5 21 24 38 47 53 57 59 74 78 88 89 99
*/
