/*

Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку

Sample Input:

zaabcbd
Sample Output:

zcd

*/

package main

import "fmt"

func main() {
	var st string
	fmt.Scan(&st)
	for i := 0; i < len(st); i++ {
		n := 1
		for j := 0; j < len(st); j++ {
			if st[i] == st[j] && i != j {
				n = 0
				break
			}
		}
		if n == 1 {
			fmt.Print(string(st[i]))
		}
	}
}
