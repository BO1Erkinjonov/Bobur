/*
Напишите программу, принимающая на вход число

N(N≥4), а затем N целых чисел-элементов среза. На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.

Sample Input:

5
41 -231 24 49 6
Sample Output:

49

*/

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var nums []int
	for i := 0; i < n; i++ {
		el := 0
		fmt.Scan(&el)
		nums = append(nums, el)
	}
	fmt.Println(nums[3])
}
