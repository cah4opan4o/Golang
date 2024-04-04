package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Функция для нахождения наибольшего числа на K-ом месте
func findKthLargest(nums []int, k int) int {
	kLargest := make([]int, 0, k)

	for _, num := range nums {
		if len(kLargest) < k || num > kLargest[0] {
			kLargest = append([]int{num}, kLargest...)
			if len(kLargest) > k {
				kLargest = kLargest[:k]
			}
		}
	}

	return kLargest[len(kLargest)-1]
}

func main() {
	var size int
	fmt.Print("Введите размер массива: ")
	fmt.Scanln(&size)

	nums := make([]int, size)

	var input string
	fmt.Print("Введите элементы массива через пробел: ")
	fmt.Scanln(&input)

	// Разбиваем введенную строку на числа и заполняем массив
	numStrings := strings.Fields(input)
	for i, numStr := range numStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Ошибка:", err)
			return
		}
		nums[i] = num
	}

	var k int
	fmt.Print("Введите номер k-го наибольшего элемента: ")
	fmt.Scanln(&k)
	kthLargest := findKthLargest(nums, k)
	fmt.Printf("%d-й наибольший элемент в массиве: %d\n", k, kthLargest)
}
