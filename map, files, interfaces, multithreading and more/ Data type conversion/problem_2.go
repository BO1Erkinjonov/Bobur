/*

Представьте что вы работаете в большой компании где используется модульная архитектура. Ваш коллега написал модуль
с какой-то логикой (вы не знаете) и передает в вашу программу какие-то данные. Вы же пишете функцию которая считывает
две переменных типа string ,  а возвращает число типа int64 которое нужно получить сложением этих строк.


Но не было бы так все просто, ведь ваш коллега не пишет на Go, и он зол из-за того, что гоферам платят больше.
Поэтому он решил подшутить над вами и подсунул вам свинью. Он придумал вставлять мусор в строки перед тем как вызывать вашу функцию.


Поэтому предварительно вам нужно убрать из них мусор и конвертировать в числа. Под мусором имеются ввиду лишние символы и
спец знаки. Разрешается использовать только определенные пакеты: fmt, strconv, unicode, strings,  bytes. Они уже
импортированы, вам ничего импортировать не нужно!

Считывать и выводить ничего не нужно!

Ваша функция должна называться adding() !

Sample Input:

%^80 hhhhh20&&&&nd
Sample Output:

100


*/

package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var st1, st2 string
	fmt.Scan(&st1, &st2)
	fmt.Println(adding(st1, st2))
}

func adding(str1, str2 string) int64 {
	cleanStr1 := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, str1)
	cleanStr2 := strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, str2)

	num1, _ := strconv.ParseInt(cleanStr1, 10, 64)
	num2, _ := strconv.ParseInt(cleanStr2, 10, 64)

	return num1 + num2
}
