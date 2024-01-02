/*

На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).

Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.

Sample Input:

13.03.2018 14:00:15,12.03.2018 14:00:15
Sample Output:

24h0m0s

*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	tim, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	time1 := strings.Split(tim, ",")

	ptime1, err := time.Parse("02.01.2006 15:04:05", time1[0][:19])
	if err != nil {
		fmt.Println(err)
		return
	}
	ptime2, err := time.Parse("02.01.2006 15:04:05", time1[1][:19])
	if err != nil {
		fmt.Println(err)
		return
	}
	var dur time.Duration
	if ptime1.After(ptime2) {
		dur = ptime1.Sub(ptime2)
	} else {
		dur = ptime2.Sub(ptime1)
	}

	fmt.Println(dur)

}
