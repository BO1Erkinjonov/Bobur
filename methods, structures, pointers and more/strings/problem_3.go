/*

Даются две строки X и S. Нужно найти и вывести первое вхождение подстроки S в строке X. Если подстроки S нет в строке X - вывести -1

Sample Input:

awesome
es
Sample Output:

2

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var x1, x2 string
	fmt.Scan(&x1, &x2)
	n := strings.Index(x1, x2)
	fmt.Println(n)
}
