/*

Напишите функцию которая принимает канал и число N типа int. Необходимо вернуть значение N+1 в канал.
Функция должна называться task().

*/

package main

import "fmt"

func main() {
	ch := make(chan int)
	var n int
	fmt.Scan(&n)
	go task(ch, n)
	fmt.Println(<-ch)
}

func task(n chan int, N int) {
	n <- N + 1
}
