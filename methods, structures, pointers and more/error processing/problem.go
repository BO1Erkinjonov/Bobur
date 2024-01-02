/*

Вы должны вызвать функцию divide которая делит два числа, но она возвращает не только результат деления, но и информацию об ошибке.
В случае какой-либо ошибки вы должны вывести "ошибка", если ошибки нет - результат функции. Функция divide(a int, b int) (int, error)
получает на вход два числа которые нужно поделить и возвращает результат (int) и ошибку (error).

Не забудьте считать два числа которые необходимо поделить (передать в функцию)!

Sample Input:

10 5
Sample Output:

2

*/

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	_, el := divide(a, b)
	if el != nil {
		fmt.Println("ошибка")
	} else {
		fmt.Println(a / b)
	}
}

func divide(a int, b int) (int, error) {
	if a < b || b == 0 {
		return 0, fmt.Errorf("error")
	}
	return a / b, nil
}
