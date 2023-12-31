/*

Дано трехзначное число. Переверните его, а затем выведите.

Формат входных данных
На вход дается трехзначное число, не оканчивающееся на ноль.

Формат выходных данных
Выведите перевернутое число.

Sample Input:

843
Sample Output:

348

*/

package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n)
	for n != 0 {
		m *= 10
		m += n % 10
		n /= 10
	}
	fmt.Println(m)
}
