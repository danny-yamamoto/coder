package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bf := bufio.NewReader(os.Stdin)
	rs, _ := bf.ReadString('\n')
	in := strings.TrimSpace(rs)
	nst := strings.Split(in, " ")
	n, _ := strconv.Atoi(nst[0])
	s, _ := strconv.Atoi(nst[1])
	t, _ := strconv.Atoi(nst[2])
	reminder := n % 2
	if t == 0 {
		// 最後の石を取り除いた人が勝ち
		// 後攻の勝ち
		if reminder == 0 {
			if s == 0 {
				// 偶数
				// Alice 先攻
				fmt.Println("Bob")
			} else {
				// 奇数
				// Bob 先攻
				fmt.Println("Alice")
			}
		} else {
			// 奇数
			// 先攻の勝ち
			if s == 0 {
				// Alice 先攻
				fmt.Println("Alice")
			} else {
				// Bob 先攻
				fmt.Println("Bob")
			}
		}
	} else {
		// 最後の石を取り除いた人が負け
		if reminder == 0 {
			// 偶数
			// 先攻の勝ち
			if s == 0 {
				fmt.Println("Alice")
			} else {
				fmt.Println("Bob")
			}
		} else {
			// 奇数
			// 先攻の負け
			if s == 0 {
				fmt.Println("Bob")
			} else {
				fmt.Println("Alice")
			}
		}
	}
}
