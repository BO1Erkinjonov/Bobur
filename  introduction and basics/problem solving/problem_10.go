/*

Найдите самое большее число на отрезке от a до b, кратное 7 .

Входные данные
Вводится два целых числа a и b (a≤b).

Выходные данные
Найдите самое большее число на отрезке от a до b (отрезок включает в себя числа a и b), кратное 7 , или выведите "NO" - если таковых нет.

Sample Input:

100
500
Sample Output:

497
*/

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for b%7 != 0 && b >= a {
		b--
	}
	if b%7 != 0 || b < a {
		fmt.Println("NO")
	} else {
		fmt.Println(b)
	}
}
