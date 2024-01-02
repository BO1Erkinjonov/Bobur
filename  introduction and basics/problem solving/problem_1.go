/*

Дано трехзначное число. Найдите сумму его цифр.

Формат входных данных
На вход дается трехзначное число.

Формат выходных данных
Выведите одно целое число - сумму цифр введенного числа.

Sample Input:

745
Sample Output:

16
*/

package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scan(&n)
	for n != 0 {
		sum += n % 10
		n /= 10
	}
	fmt.Println(sum)
}
