/*

Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельных горутинах вызвать функцию work1() 10 раз и дождаться результатов выполнения вызванных функций.


Функция work1() ничего не принимает и не возвращает.

*/

package main

import (
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			work1()
			wg.Done()
		}(wg)
	}
	wg.Wait()
}

func work1() {
	time.Sleep(time.Second)
}
