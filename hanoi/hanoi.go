package main

import "fmt"

func main() {
	n := 3
	x, y, z := "x", "y", "z"
	hanoi(n, x, y, z)
}

func hanoi(n int, x, y, z string) {
	if n == 0 {
		return
	} else {
		hanoi(n-1, x, z, y)            // 将 n-1个盘从 x 借助y移到 z
		fmt.Printf("%s -> %s\n", x, y) // 将当前x最下面的一个盘移到y
		hanoi(n-1, z, y, x)            // 将 n-1个盘从 z 借助x移到 y
	}
}
