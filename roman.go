package main

import "fmt"

func dispRep(num int, text string) string {
	result := ""
	for i := 0; i < num; i++ {
		result += text
	}
	return result
}

func prefix(m int) string {
	resulta := ""
	switch {
	case m == 10:
		resulta = "C"
	case m == 9:
		resulta = "XC"
	case m >= 5:
		resulta = "L" + dispRep(m-5, "X")
	case m == 4:
		resulta = "XL"
	default:
		resulta = dispRep(m, "X")
	}
	return resulta
}

func suffix(n int) string {
	resultb := ""
	switch {
	case n == 9:
		resultb = "IX"
	case n >= 5:
		resultb = "V" + dispRep(n-5, "I")
	case n == 4:
		resultb = "IV"
	default:
		resultb = dispRep(n, "I")
	}
	return resultb
}

func main() {
	// write GO to return ROMAN number from numeric numbers
	for i := 1; i <= 100; i++ {
		result := prefix(i/10) + suffix(i%10)
		fmt.Printf("%d = %s \n", i, result)
	}
}
