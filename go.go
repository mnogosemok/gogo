package main

import (
	"fmt"
	"strconv"
	"strings"
)

const op = "+-*/"

func main() {
	var x, a, y string
	fmt.Scan(&x, &a, &y)
	fmt.Println(result(x, a, y))
}

func result(x, a, y string) string {
	parsSign(a)
	first, err1 := strconv.Atoi(x)
	second, err2 := strconv.Atoi(y)
	if err1 == nil && err2 == nil {
		return strconv.Itoa(parsNumber(first, second, a))
	} else {
		return parsRome(x, a, y)
	}
}

func parsSign(a string) string {
	if strings.ContainsAny(a, op) {
		return a
	}
	panic("На вход поступил неверный арифметический знак")
}

func parsNumber(x, y int, a string) int {
	var res int
	if x < 11 && y < 11 && x > 0 && y > 0 {
		switch a {
		case "+":
			res = x + y
		case "-":
			res = x - y
		case "*":
			res = x * y
		case "/":
			res = x / y
		}
		return res
	} else {
		panic("Нужны только арбские либо только рисмские числа не больше 10 и не меньше 1")
	}

}

func parsRome(x, a, y string) string {
	numRome := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	}
	firstRome, ok1 := numRome[x]
	secondRome, ok2 := numRome[y]
	if ok1 && ok2 {
		num := parsNumber(firstRome, secondRome, a)
		if num < 0 {
			panic("В римской системе нет отрицательных чисел.")
		}
		symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
		values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
		var result strings.Builder
		for i := 0; i < len(symbols); i++ {
			for num >= values[i] {
				result.WriteString(symbols[i])
				num -= values[i]
			}
		}
		return result.String()
	} else {
		panic("Нужны только арбские либо только рисмские числа не больше 10")
	}
}
