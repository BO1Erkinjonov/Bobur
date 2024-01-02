/*

Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.

Пример вызова вашей функции:

a, b := sumInt(1, 0)
fmt.Println(a, b)

Результат: 2, 1

*/

package main

import "fmt"

func main() {
	var n, num int
	var sl []int
	fmt.Print("Nechta son kiritasiz: ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		sl = append(sl, num)
	}
	a, b := sumInt(sl...)
	fmt.Printf("Uzunligi: %d\nYegidisi: %d", a, b)
}
func sumInt(args ...int) (int, int) {
	sum := 0
	for _, num := range args {
		sum += num
	}
	return len(args), sum
}
