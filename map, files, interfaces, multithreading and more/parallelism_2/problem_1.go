/*

Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельной горутине вызвать функцию work() и дождаться результатов ее выполнения.

Функция work() ничего не принимает и не возвращает.

*/

package main

import "time"

func main() {
	done := make(chan struct{})
	go func() {
		work()
		close(done)
	}()
	<-done

}

func work() {
	time.Sleep(time.Second)
}
