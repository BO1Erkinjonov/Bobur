/*

В ходе анализа результатов переписи населения информация была сохранена в объекте типа map:

groupCity := map[int][]string{
	10:   []string{...}, // города с населением 10-99 тыс. человек
	100:  []string{...}, // города с населением 100-999 тыс. человек
	1000: []string{...}, // города с населением 1000 тыс. человек и более
}
При подготовке важного отчета о городах с населением 100-999 тыс. человек был подготовлен другой объект типа map:

cityPopulation := map[string]int{
	город: население города в тысячах человек,
}
Однако база данных с информацией о точной численности населения содержала ошибки, поэтому в cityPopulation в т.ч. была
сохранена информация о городах, которые входят в другие группы из groupCity.

Ваша программа имеет доступ к обоим указанным отображениям, требуется исправить cityPopulation, чтобы в ней была сохранена
информация только о городах из группы groupCity[100].

Функция main() уже объявлена, доступ к отображениям осуществляется по указанным именам. По результатам выполнения ничего
больше делать не требуется, проверка будет осуществлена автоматически.

*/

package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   []string{"Село", "Деревня", "ПГТ"},
		100:  []string{"Город", "Станица"},
		1000: []string{"Мегаполис", "Миллионник"},
	}

	cityPopulation := map[string]int{
		"Город":     101,
		"Станица":   102,
		"Село":      103,
		"Мегаполис": 104,
	}
	for _, in := range append(groupCity[10], groupCity[1000]...) {
		delete(cityPopulation, in)
	}
	fmt.Println(groupCity)
}
