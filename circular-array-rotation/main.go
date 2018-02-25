// HackerRank problem
// https://www.hackerrank.com/challenges/circular-array-rotation/problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swap(s *[]int) {
	l := len((*s)) - 1
	(*s) = append([]int{(*s)[l]}, (*s)[:l]...)
}

func main() {
	cnt := 0
	var n, k, q int
	var seq []int
	var m []int

	scn := bufio.NewScanner(os.Stdin)
	for scn.Scan() {
		switch cnt {
		case 0:
			tmp := strings.Split(scn.Text(), " ")
			n, _ = strconv.Atoi(tmp[0])
			k, _ = strconv.Atoi(tmp[1])
			q, _ = strconv.Atoi(tmp[2])
		case 1:
			tmp := strings.Split(scn.Text(), " ")
			if len(tmp) != n {
				fmt.Printf("wrong seq length %v, %v", len(tmp), n)
			}
			seq = make([]int, 0, n)
			m = make([]int, 0, n)
			for _, v := range tmp {
				val, _ := strconv.Atoi(v)
				seq = append(seq, val)
			}
		default:
			tmp, _ := strconv.Atoi(scn.Text())
			m = append(m, tmp)
		}

		cnt++
	}
	if len(m) != q {
		fmt.Printf("wrong m length %v, %v", len(m), q)
	}
	// make rotaions
	for i := 0; i < k; i++ {
		swap(&seq)
	}

	for _, i := range m {
		fmt.Println(seq[i])
	}
}

// =============== this work with full input.
// package main

// import "fmt"

// func CircularArrayRotation() {
//     var n, k, q int

//     fmt.Scanf("%d%d%d", &n, &k, &q)
//     a := make([]int, n)
//     for i := range a {
//         fmt.Scanf("%d", &a[i])
//     }

//     var i int
//     for t := 0; t < q; t++ {

//         fmt.Scanf("%d", &i)
//         j := i - k
//         for j < 0 {
//             j += n
//         }

//         fmt.Println(a[j])
//     }

// }

// func main() {
//     //Enter your code here. Read input from STDIN. Print output to STDOUT
//     CircularArrayRotation()
// }
