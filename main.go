package main

import (
	"fmt"
)

func remove(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 5)
	slice3 := make([]int, 0, 5)

	fmt.Println("Инициализация слайсов:")
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	fmt.Println("slice3:", slice3)

	slice1 = append(slice1, 4, 5)
	fmt.Println("\nC помощью append в slice1 добавляем два элемента:", slice1)

	slice1 = remove(slice1, 2)
	fmt.Println("\nУдаляем из slice1 элемент по индексу 2:", slice1)

	array := [5]int{1, 2, 3, 4, 5}
	slice4 := array[1:4]
	fmt.Println("\nИз массива array", array, "делаем срез slice4", slice4)

	slice4[0] = 90
	fmt.Println("\nИзменение 1 элемента массива slice4", slice4)

}
