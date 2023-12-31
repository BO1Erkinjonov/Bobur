/*

На вход подаются a и b - катеты прямоугольного треугольника. Нужно найти длину гипотенузы

Sample Input:

6 8
Sample Output:

10

*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	fmt.Scan(&a, &b)
	c = a*a + b*b
	fmt.Println(math.Sqrt(float64(c)))
}
