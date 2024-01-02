/*

Вам необходимо написать функцию calculator следующего вида:

 "func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int"

Функция получает в качестве аргументов 3 канала, и возвращает канал типа <-chan int.

в случае, если аргумент будет получен из канала firstChan, в выходной (возвращенный) канал вы должны отправить квадрат
аргумента.

в случае, если аргумент будет получен из канала secondChan, в выходной (возвращенный) канал вы должны отправить
результат умножения аргумента на 3.

в случае, если аргумент будет получен из канала stopChan, нужно просто завершить работу функции.

Функция calculator должна быть неблокирующей, сразу возвращая управление. Ваша функция получит всего одно значение в один из
каналов - получили значение, обработали его, завершили работу.

После завершения работы необходимо освободить ресурсы, закрыв выходной канал, если вы этого не сделаете, то превысите
предельное время работы.

*/

package main

import "fmt"

func main() {
	type баста struct{}
	fir, sec := make(chan int), make(chan int)
	stop := make(chan struct{})
	go func() {
		// fir <- 5
		// close(fir)
		sec <- 4
		stop <- баста{}
	}()
	fmt.Print(<-calculator(fir, sec, stop))
}

func calculator(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {
	sum := make(chan int)
	go func() {
		defer close(sum)
		select {
		case n := <-firstChan:
			sum <- n * n
		case n := <-secondChan:
			sum <- n * 3
		case <-stopChan:
			return
		}
	}()
	return sum
}
