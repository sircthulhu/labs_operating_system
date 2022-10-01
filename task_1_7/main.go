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

// DecryptBlock дешифрует блок под номером number в массиве blocks с помощью ключа key
func DecryptBlock(blocks *[]byte, key *[]byte, number int) {
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

	j := 0 // j для обхода ключа
	for i := blockStartIdx; i < lastIdx; i++ {
		(*blocks)[i] = (*blocks)[i] ^ (*key)[j]
		j++
	}

	// Обмениваем байты в блоке
	SwapBytes(blocks, blockStartIdx, lastIdx)
}

func main() {
	key := "111"
	encrypted := []byte{0xe0, 0x8e, 0xd0, 0x89, 0xe1, 0x80, 0xe1, 0x83, 0xd0, 0xb3, 0xe0, 0xb5}

	keyBytes := []byte(key)

	fmt.Println("Начинаю дешифровку...")
	blocksLen := len(encrypted)
	keyLen := len(key)
	blocksCount := blocksLen / keyLen

	wg := sync.WaitGroup{}
	for i := 0; i < blocksCount; i++ {
		wg.Add(1)
		go func(blockNumber int) {
			defer wg.Done()
			DecryptBlock(&encrypted, &keyBytes, blockNumber)
		}(i)
	}
	wg.Wait()

	fmt.Printf("Зашифрованная строка: %#v\n", string(encrypted))
}

/*
Тестовый запуск:
Начинаю дешифровку...
Зашифрованная строка: "Тестовая строка для шифровки"
*/
