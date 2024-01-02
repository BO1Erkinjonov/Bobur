/*

На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)

Sample Input:

ihgewlqlkot
Sample Output:

hello

*/

package main

import (
	"fmt"
)

func main() {
	var st string
	fmt.Scan(&st)
	for i := 0; i < len(st); i++ {
		if i%2 != 0 {
			fmt.Print(string(st[i]))
		}
	}
}
