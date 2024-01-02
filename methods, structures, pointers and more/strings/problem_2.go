/*
На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.

Палиндром — буквосочетание, слово или текст, одинаково читающееся в обоих направлениях (например, "topot", "oko", "zakaz").

Sample Input:

topot
Sample Output:

Палиндром
*/

package main

import "fmt"

func reverse(st string) (result string) {
	for _, v := range st {
		result = string(v) + result
	}
	return
}

func main() {
	var st string
	fmt.Scan(&st)
	st1 := reverse(st)
	if st1 == st {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
