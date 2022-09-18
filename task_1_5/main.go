package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"time"
)

// Node это узел бинарного дерева с целыми числами
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// InsertNode вставляет в дерево root новый элемент elem
// с помощью алгоритма двоичного поиска со вставкой
func InsertNode(root **Node, elem int) {
	if *root == nil {
		*root = &Node{
			Data:  elem,
			Left:  nil,
			Right: nil,
		}
	} else if (*root).Data > elem {
		InsertNode(&((*root).Left), elem)
	} else if (*root).Data < elem {
		InsertNode(&((*root).Right), elem)
	}
	// Если такой элемент уже существует, пропускаем его
}

// PrintTree выводит двоичное дерево в консоль по уровням
func PrintTree(root *Node, level int) {
	if root != nil {
		indent := strings.Repeat(" ", level*4)
		fmt.Print(indent, root.Data)
		fmt.Println()
		
		PrintTree(root.Left, level+1)
		PrintTree(root.Right, level+1)
	}
}

// CalculateNodes считает количество узлов
func CalculateNodes(root *Node) int {
	if root == nil {
		return 1
	} else {
		return CalculateNodes(root.Left) + CalculateNodes(root.Right)
	}
}

// SumDataInNodes считает сумму элементов во всех узлах
func SumDataInNodes(root *Node) int {
	if root == nil {
		return 0
	} else {
		return root.Data + CalculateNodes(root.Left) + CalculateNodes(root.Right)
	}
}

func main() {
	var count int = 15
	var tree *Node = nil
	rand.Seed(time.Now().Unix())
	
	// Заполняем дерево случайными числами
	for i := 0; i < count; i++ {
		number := rand.Intn(100)
		if rand.Float32() > 0.5 {
			number = -number
		}
		
		InsertNode(&tree, number)
	}
	fmt.Println("Созданное дерево:")
	PrintTree(tree, 0)
	
	wg := sync.WaitGroup{}
	// Вычисляем количество узлов
	countLeft := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		countLeft = CalculateNodes(tree.Left)
	}()
	wg.Add(1)
	countRight := 0
	go func() {
		defer wg.Done()
		countRight = CalculateNodes(tree.Right)
	}()
	
	// Вычисляем сумму чисел
	wg.Add(1)
	sumLeft := 0
	go func() {
		defer wg.Done()
		sumLeft = SumDataInNodes(tree.Left)
	}()
	
	wg.Add(1)
	sumRight := 0
	go func() {
		defer wg.Done()
		sumRight = SumDataInNodes(tree.Right)
	}()
	
	wg.Wait()
	sum := sumLeft + sumRight
	count = countLeft + countRight
	
	fmt.Printf("Сумма элементов: %d. Количество элементов: %d. Среднее: %f", sum, count, float64(sum)/float64(count))
}

/*
Тестовый запуск:
Созданное дерево:
60
    -57
        -85
            -80
        -10
            -36
            8
                1
                    -6
                46
                    27
    95
        75
            69
            89
Сумма элементов: 54. Количество элементов: 16. Среднее: 3.375000
Process finished with the exit code 0

*/
