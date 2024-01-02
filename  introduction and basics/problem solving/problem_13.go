/*

Номер числа Фибоначчи
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φn=A. Если А не является числом Фибоначчи, выведите число -1.

Входные данные

Вводится натуральное число.

Выходные данные

Выведите ответ на задачу.

Sample Input:

8
Sample Output:

6

*/

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n == 0 {
		fmt.Println(1)
	} else if n == 1 {
		fmt.Println(2)
	} else {
		i := 2
		j := 1
		s := 3
		for i < n {
			i += j
			j = i - j
			s++
		}
		if i == n {
			fmt.Println(s)
		} else {
			fmt.Println(-1)
		}
	}
}
