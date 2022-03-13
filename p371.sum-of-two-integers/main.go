package main

func getSum(a int, b int) int {
	for b != 0 {
		sum := a ^ b
		carry := (a & b)
		//fmt.Printf("a=%064b\n", uint(a))
		//fmt.Printf("b=%064b\n", uint(b))
		//fmt.Printf("c=%064b\n", uint(sum))
		//fmt.Printf("d=%064b\n\n", uint(carry))
		a = sum
		b = carry << 1
	}
	return a
}
