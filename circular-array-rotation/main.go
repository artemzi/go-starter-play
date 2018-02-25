// HackerRank problem
// https://www.hackerrank.com/challenges/circular-array-rotation/problem
package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	cnt := 0
	n, k, q := 0, 0, 0
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
				panic("wrong sequence length")
			}
			seq = make([]int, 0, n)
			for _, v := range tmp {
				val, _ := strconv.Atoi(v)
				seq = append(seq, val)
			}
		default:
			tmp, _ := strconv.Atoi(scn.Text())
			m = make([]int, 0, n-1)
			m = append(m, tmp)
		}

		cnt++
	}
	log.Printf("\nn = %d, k = %d, q = %d\nseq = %d\nm = %d\n", n, k, q, seq, m)
}
