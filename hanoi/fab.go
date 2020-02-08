package main

func main() {
	println(fb(11))
}

// n 是天数
func fb(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fb(n-1) + fb(n-2)
}
