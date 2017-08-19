package main

import (
	"os"
	"fmt"
	"regexp"
	"math"
	"strconv"
)

const (
	pattern = `\d+\.\d+|[+\-*/^]|\d+`
	count   = -1
)

func main() {

	switch len(os.Args) {
	case 1:
		fmt.Println("You must entered an expression.")
	case 2:
		parse(os.Args[1])
	default:
		fmt.Println("Too many arguments.")
	}
}
func parse(expression string) {
	prepare(regexp.MustCompile(pattern).FindAllString(expression, count))
}
func prepare(args []string) {
	if args != nil && len(args) == 3 {
		arg1, err1 := strconv.ParseFloat(args[0], 64)
		arg2, err2 := strconv.ParseFloat(args[2], 64)
		if err1 == nil && err2 == nil{
			calc(arg1, arg2, args[1])
		}else{
			fmt.Println("Wrong expression.")
		}
	} else {
		fmt.Println("Wrong expression.")
	}
}
func calc(arg1 float64, arg2 float64, sign string) {
	switch sign {
	case "+":
		fmt.Println(arg1 + arg2)
	case "-":
		fmt.Println(arg1 - arg2)
	case "*":
		fmt.Println(arg1 * arg2)
	case "/":
		fmt.Println(arg1 / arg2)
	case "^":
		fmt.Println(math.Pow(arg1, arg2))
	default:
		fmt.Println("Unknown sign ", sign)
	}
}
