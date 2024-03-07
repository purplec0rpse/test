package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func rometoarab(a string) int {
	var x int
	switch rome := a; rome {
	case "I":
		x = 1
	case "II":
		x = 2
	case "III":
		x = 3
	case "IV":
		x = 4
	case "V":
		x = 5
	case "VI":
		x = 6
	case "VII":
		x = 7
	case "VIII":
		x = 8
	case "IX":
		x = 9
	case "X":
		x = 10
	default:
		panic(a)
	}
	return x
}

func arabtorome(a int) string {
	var x, x1, x10 string
	var (
		table1  = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
		table10 = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC", "C"}
	)
	x10 = table10[a/10]
	x1 = table1[a%10]
	x = x10 + x1
	return x
}

func operation(x int, c string, y int) int {
	var z int
	switch oper := c; oper {
	case "+":
		z = x + y
	case "-":
		z = x - y
	case "*":
		z = x * y
	case "/":
		z = x / y
	default:
		panic(c)
	}
	return z
}

func convert(a string) int {
	x, e := strconv.Atoi(a)
	if e != nil {
		panic(e)
	}
	return x
}

func main() {
	var a, b, s string
	var arome, brome bool
	var x, y int
	var re = regexp.MustCompile(`I|V|X`)
	fmt.Scan(&a, &s, &b)
	if re.MatchString(a) {
		arome = true
	} else {
		arome = false
	}
	if re.MatchString(b) {
		brome = true
	} else {
		brome = false
	}
	if (arome == false) && (brome == false) {
		x = convert(a)
		y = convert(b)
		fmt.Print(operation(x, s, y))
	} else if (arome == true) && (brome == true) {
		x = rometoarab(a)
		y = rometoarab(b)
		fmt.Print(arabtorome(operation(x, s, y)))
	} else {
		panic(arome)
	}
}
