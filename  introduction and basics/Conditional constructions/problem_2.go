// По данному трехзначному числу определите, все ли его цифры различны.
//
// Формат входных данных
// На вход подается одно натуральное трехзначное число.
//
// Формат выходных данных
// Выведите "YES", если все цифры числа различны, в противном случае - "NO".
//
// Sample Input 1:
//
// 237
// Sample Output 1:
//
// YES
// Sample Input 2:
//
// 117
// Sample Output 2:
//
// NO
package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)
	a := num % 10
	num /= 10
	b := num % 10
	num /= 10
	if a != b && a != num && b != num {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
