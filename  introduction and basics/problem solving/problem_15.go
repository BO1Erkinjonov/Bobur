/*

Из натурального числа удалить заданную цифру.

Входные данные

Вводятся натуральное число и цифра, которую нужно удалить.

Выходные данные

Вывести число без заданных цифр.

Sample Input:

38012732
3
Sample Output:

801272

*/

package main

import "fmt"

func main() {
	var n1, n2 int
	var l []int
	fmt.Scan(&n1)
	fmt.Scan(&n2)
	for n1 != 0 {
		if n1%10 != n2 {
			l = append(l, n1%10)
		}
		n1 /= 10
	}
	n1 = 0
	i := len(l) - 1
	for i > -1 {
		n1 *= 10
		n1 += l[i]
		i--
	}
	fmt.Println(n1)
}
