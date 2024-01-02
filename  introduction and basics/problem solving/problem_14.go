/*
Двоичная запись
Дано натуральное число N. Выведите его представление в двоичном виде.

# Входные данные

# Задано единственное число N

# Выходные данные

Необходимо вывести требуемое представление числа N.

Sample Input:

12
Sample Output:

1100
*/
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var s []int
	for n > 0 {
		s = append(s, n%2)
		n /= 2
	}
	for i := len(s) - 1; i > -1; i-- {
		n *= 10
		n += s[i]
	}
	fmt.Println(n)
}
