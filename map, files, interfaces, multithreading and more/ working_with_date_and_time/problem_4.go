/*

На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере). Кроме того, вам дана дата в
формате Unix-Time: 1589570165 в виде константы типа int64 (наносекунды для целей преобразования равны 0).

Требуется считать данные о продолжительности периода, преобразовать их в тип Duration, а затем вывести (в формате UnixDate)
дату и время, получившиеся при добавлении периода к стандартной дате.

Небольшая подсказка: базовую дату необходимо явно перенести в зону UTC с помощью одноименного метода.

Sample Input:

12 мин. 13 сек.
Sample Output:

Fri May 15 19:28:18 UTC 2020

*/

package main

import (
	"fmt"
	"time"
)

const now = 1589570165

func main() {
	var m, s int
	fmt.Scanf("%d мин. %d сек.", &m, &s)
	dur := time.Minute*time.Duration(m) + time.Second*time.Duration(s)
	basetime := time.Unix(now, 0).UTC()
	unix1 := basetime.Add(dur)
	fmt.Println(unix1.Format(time.UnixDate))
}
