/*

Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

Входные данные

Вводится четыре числа.

Выходные данные

Необходимо вернуть из функции наименьшее из 4-х данных чисел

Sample Input:

4 5 6 7
Sample Output:

4

*/

package main

import "fmt"

func main() {
	fmt.Println(minimumFromFour())
}

func minimumFromFour() int {
	var min int
	for i := 0; i < 4; i++ {
		var n int
		fmt.Scan(&n)
		if i == 0 {
			min = n
		} else {
			if min > n {
				min = n
			}
		}
	}
	return min
}