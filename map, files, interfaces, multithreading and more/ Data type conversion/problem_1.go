/*

Напишите функцию с именем convert, которая конвертирует входное число типа int64 в uint16.

Считывать и выводить ничего не нужно!

*/

package main

import "fmt"

func main() {
	var n int64
	num := convert(n)
	fmt.Printf("%T", num)
}

func convert(x int64) uint16 {
	return uint16(x)
}
