/*

Найдите количество минимальных элементов в последовательности.

Входные данные

Вводится натуральное число N, а затем N целых чисел последовательности.

Выходные данные

Выведите количество минимальных элементов последовательности.

Sample Input:

3
21
11
4
Sample Output:

1


*/

package main

import "fmt"

func main() {
	var n, n1, min int
	j := 1
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&n1)
		if i == 0 {
			min = n1
		} else if min > n1 {
			min = n1
			j = 1
		} else if min == n1 {
			j++
		}
	}
	fmt.Println(j)
}
