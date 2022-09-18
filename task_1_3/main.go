package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
				(*col)[j] = rand.Intn(99) + 1 // Числа должны быть натуральными
			}
		}(&matrix[col])
	}
	wg.Wait()
	
	return &matrix
}

// multiplyRow умножает конкретную строку row матрицы matrix на multiplier
func multiplyRow(row int, matrix *[][]int, multiplier int) {
	cols := len((*matrix)[row])
	for i := 0; i < cols; i++ {
		(*matrix)[row][i] *= multiplier
	}
}

func main() {
	n := 6
	m := 8
	var multiplier int
	
	// matrix это двумерный массив целых чисел
	matrix := createRandomMatrix(n, m)
	
	fmt.Println("Созданная матрица:")
	printMatrix(matrix)
	
	fmt.Print("\nВведите множитель: ")
	r := bufio.NewReader(os.Stdin)
	multiplierString, err := r.ReadString('\n')
	if err != nil {
		fmt.Printf("Ошибка считывания множителя: %#v\n", err)
		os.Exit(1)
	}
	multiplier, err = strconv.Atoi(strings.Trim(multiplierString, "\n"))
	if err != nil {
		fmt.Printf("Неверный формат числа множителя. Ошибка: %#v\n", err)
		os.Exit(1)
	}
	
	fmt.Println("\nУмножаем матрицу на множитель...")
	
	wg := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(row int) {
			defer wg.Done()
			multiplyRow(row, matrix, multiplier)
		}(i)
	}
	wg.Wait()
	fmt.Println("Результирующая матрица:")
	printMatrix(matrix)
}

/*
Тестовый запуск:
Созданная матрица:
62 95 83 82 47 82 17 56
95 7 60 89 22 12 68 18
19 39 75 6 75 1 33 11
10 89 17 29 25 76 23 52
6 24 31 45 79 62 13 82
24 79 21 48 14 16 8 18

Введите множитель: 2

Умножаем матрицу на множитель...
Результирующая матрица:
124 190 166 164 94 164 34 112
190 14 120 178 44 24 136 36
38 78 150 12 150 2 66 22
20 178 34 58 50 152 46 104
12 48 62 90 158 124 26 164
48 158 42 96 28 32 16 36

Process finished with the exit code 0
*/
