package main

import (
	"reflect"
	"testing"
)

func TestSwapBytes(t *testing.T) {
	t.Parallel()
	t.Run("Размер нечетный и блок полностью", func(t *testing.T) {
		t.Parallel()
		block := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
		expectedBlock := []byte{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		
		SwapBytes(&block, 0, len(block)-1)
		if !reflect.DeepEqual(block, expectedBlock) {
			t.Errorf("Блок неверным образом заменился.\nОжидаемое: %v\nПолученное:%v\n", expectedBlock, block)
		}
	})
	
	t.Run("Размер четный и блок полностью", func(t *testing.T) {
		t.Parallel()
		block := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
		expectedBlock := []byte{14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		
		SwapBytes(&block, 0, len(block)-1)
		if !reflect.DeepEqual(block, expectedBlock) {
			t.Errorf("Блок неверным образом заменился.\nОжидаемое: %v\nПолученное:%v\n", expectedBlock, block)
		}
	})
	
	t.Run("Размер больше startIdx и endIdx и четный", func(t *testing.T) {
		t.Parallel()
		block := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
		expectedBlock := []byte{1, 2, 10, 9, 8, 7, 6, 5, 4, 3, 11, 12, 13, 14, 15}
		
		SwapBytes(&block, 2, 9)
		if !reflect.DeepEqual(block, expectedBlock) {
			t.Errorf("Блок неверным образом заменился.\nОжидаемое: %v\nПолученное:%v\n", expectedBlock, block)
		}
	})
	
	t.Run("Размер больше startIdx и endIdx и нечетный", func(t *testing.T) {
		t.Parallel()
		block := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
		expectedBlock := []byte{1, 2, 9, 8, 7, 6, 5, 4, 3, 10, 11, 12, 13, 14, 15}
		
		SwapBytes(&block, 2, 8)
		if !reflect.DeepEqual(block, expectedBlock) {
			t.Errorf("Блок неверным образом заменился.\nОжидаемое: %v\nПолученное:%v\n", expectedBlock, block)
		}
	})
}

func TestEncryptBlock(t *testing.T) {
	t.Parallel()
	t.Run("Длина строки не кратна длине ключа", func(t *testing.T) {
		str := "Строка для дешифровки определенной длины" // длина 40
		key := "Ключ шифрования"                          // длина 15
		
		strBytes := []byte(str)
		keyBytes := []byte(key)
		for i := 0; i < len(strBytes)/len(keyBytes); i++ {
			EncryptBlock(&strBytes, &keyBytes, i)
		}
		
		expected := []byte("ынилд йоннеледерпо икворфишед ялд акортС")
		j := 0
		for i := 0; i < len(expected); i++ {
			expected[i] = expected[i] ^ keyBytes[j]
			
			// Проходимся по ключу по кругу
			j++
			if j >= len(keyBytes) {
				j = 0
			}
		}
	})
}
