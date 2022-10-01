package main

import (
	"fmt"
	"sync"
)

// SwapBytes обменивает местами элементы в blocks, начиная с startIdx, заканчивая endIdx
// например, {1,2,3,4} должно превратиться в {4,3,2,1}
func SwapBytes(blocks *[]byte, startIdx, endIdx int) {
	middle := (endIdx + startIdx) / 2
	var temp byte
	i := startIdx
	j := endIdx
	for i <= middle {
		temp = (*blocks)[i]
		(*blocks)[i] = (*blocks)[j]
		(*blocks)[j] = temp

		i++
		j--
	}
}

// EncryptBlock шифрует блок под номером number в массиве blocks с помощью ключа key
func EncryptBlock(blocks *[]byte, key *[]byte, number int) {
	blocksLen := len(*blocks)
	keyLen := len(*key)
	blockStartIdx := keyLen * number
	if blockStartIdx > blocksLen {
		return
	}

	lastIdx := blockStartIdx + keyLen - 1
	if lastIdx >= blocksLen {
		lastIdx = blocksLen - 1
	}

	// Обмениваем байты в блоке
	SwapBytes(blocks, blockStartIdx, lastIdx)
	// Шифруем с помощью ключа
	j := 0 // j для обхода ключа
	for i := blockStartIdx; i < lastIdx; i++ {
		(*blocks)[i] = (*blocks)[i] ^ (*key)[j]
		j++
	}
}

func main() {
	str := "привет"
	fmt.Printf("Исходная строка: %s\n", str)
	key := "111"
	fmt.Printf("Ключ для шифрования текста: %s\n", key)

	fmt.Println("Начинаю шифрование...")
	blocksLen := len(str)
	keyLen := len(key)
	blocksCount := blocksLen / keyLen

	blocks := []byte(str)
	keyBytes := []byte(key)
	wg := sync.WaitGroup{}
	for i := 0; i < blocksCount; i++ {
		wg.Add(1)
		go func(blockNumber int) {
			defer wg.Done()
			EncryptBlock(&blocks, &keyBytes, blockNumber)
		}(i)
	}

	wg.Wait()

	fmt.Printf("Зашифрованная строка: %#v\n", blocks)
}

/*
Тестовый запуск:
Исходная строка: Тестовая строка для шифровки
Ключ для шифрования текста: Ключ шифрования
Начинаю шифрование...
Зашифрованная строка: []byte{0xf0, 0x2a, 0x0, 0x1, 0x1, 0x30, 0x1, 0x7, 0xf1, 0x53, 0x59, 0x51, 0x69, 0xf1, 0xb, 0x0, 0x30, 0x0, 0xc, 0x0, 0xc, 0x0, 0x32, 0x1, 0x3c, 0x1, 0xd, 0x1, 0x2d, 0xd0, 0xd0, 0xb4, 0xd0, 0xbb, 0xd1, 0x8f, 0x20, 0xd1, 0x88, 0xd0, 0xb8, 0xd1, 0x84, 0xd1, 0x80, 0xd0, 0xbe, 0xd0, 0xb2, 0xd0, 0xba, 0xd0, 0xb8}

*/
