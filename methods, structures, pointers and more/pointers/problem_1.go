/*

Напишите функцию, которая умножает значения на которые ссылаются два указателя и выводит получившееся произведение в консоль. Ввод данных уже написан за вас.

Sample Input:

2 2
Sample Output:

4

*/

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	test(&a, &b)
}
func test(x1 *int, x2 *int) {
	fmt.Println(*x1 * *x2)
}
