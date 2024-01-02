/*

Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool, int, int соответственно.
У этой структуры должны быть методы: Shoot и RideBike, которые не принимают аргументов, но возвращают значение bool.

Если значение On == false, то оба метода вернут false.

Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод возвращает true),
если его нет, то метод вернет false. Метод RideBike работает также, но только зависит от свойства Power.

*/

package main

import (
	"fmt"
)

type Test_struct struct {
	On    bool
	Ammo  int
	Power int
}

func (a *Test_struct) Shoot() bool {
	if a.On == false {
		return false
	}
	if a.Ammo > 0 {
		a.Ammo--
		return true
	}

	return false
}
func (a *Test_struct) RideBike() bool {
	if a.On == false {
		return false
	}
	if a.Power > 0 {
		a.Power--
		return true
	}
	return false
}

func main() {
	testStruct := &Test_struct{false, 2, 3}
	fmt.Println(testStruct.Shoot(), testStruct.RideBike(), testStruct)
}
