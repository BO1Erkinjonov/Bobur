/*

Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку в заданный канал 5 раз, добавив к ней пробел.

Функция должна называться task2().

*/

package main

import "fmt"

func main() {
	ch := make(chan string)
	var st string
	fmt.Scan(&st)
	go task2(ch, st)
	for i := range ch {
		fmt.Println(i)
	}
}

func task2(st1 chan string, st2 string) {
	for i := 0; i < 5; i++ {
		st1 <- st2 + " "
	}

}
