/*

Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля должна быть не менее 5 символов, он должен содержать только арабские цифры и буквы латинского алфавита. На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"

Sample Input:

fdsghdfgjsdDD1
Sample Output:

Ok

*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	var st string
	fmt.Scan(&st)
	var lb string = "ABCDEFJHIGKLMNOPQRSTUVWXYZabcdefjhigklmnopqrstuvwxyz1234567890"
	if len(st) < 5 {
		fmt.Println("Wrong password")
		return
	}
	for i := 0; i < len(st); i++ {
		if !strings.Contains(lb, string(st[i])) {
			fmt.Println("Wrong password")
			return
		}
	}
	fmt.Println("Ok")
}
