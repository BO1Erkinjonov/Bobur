/*

Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.

Входные данные

Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только арабские цифры.

Выходные данные

Выведите максимальную цифру, которая встречается во введенной строке.

Sample Input:

1112221112
Sample Output:

2

*/

package main

import (
	"fmt"
)

func main() {
	var st1 string
	fmt.Scan(&st1)
	max := st1[0] - '0'
	for i := 0; i < len(st1); i++ {
		v := st1[i] - '0'
		if v > max {
			max = v
		}
	}
	fmt.Println(max)
}
