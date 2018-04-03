/*
Numeros the Artist had two lists that were permutations of one another.
He was very proud. Unfortunately, while transporting them from one exhibition to another,
some numbers were lost out of the first list. Can you find the missing numbers?

Notes

If a number occurs multiple times in the lists, you must ensure that
the frequency of that number in both lists is the same. If that is not the case,
then it is also a missing number.
You have to print all the missing numbers in ascending order.
Print each missing number once, even if it is missing multiple times.
The difference between maximum and minimum number in the second list is less than or equal to .
Input Format

There will be four lines of input:

n - the size of the first list, arr
The next line contains n space-separated integers arr
m - the size of the second list, brr
The next line contains m space-separated integers brr
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int

	arr := make([]int, 10001)
	brr := make([]int, 10001)

	io := bufio.NewReader(os.Stdin)

	fmt.Fscan(io, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(io, &m)
		arr[m]++
	}

	fmt.Fscan(io, &n)
	for i := 0; i < n; i++ {
		fmt.Fscan(io, &m)
		brr[m]++
	}

	// Missing numbers
	for i := 0; i < 10001; i++ {
		if brr[i]-arr[i] > 0 {
			fmt.Printf("%v ", i)
		}
	}
}
