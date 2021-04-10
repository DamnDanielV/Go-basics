package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

// operate: recibe una operacion matematica e imprime el resultado
func (calc) operate(entrada string) {
	nums := strings.Split(entrada, " ")
	num1 := parser(nums[0])
	num2 := parser(nums[2])
	switch nums[1] {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "/":
		fmt.Println(num1 / num2)
	case "*":
		fmt.Println(num1 * num2)
	default:
		fmt.Println("operacion incorrecta")
	}
}

// parser: recibe un numero tipo string y lo retorna en formato int
// si recibe una letra imprime un mensaje y retorna 0
func parser(numString string) int {
	operador1, err := strconv.Atoi(numString)
	if err != nil {
		fmt.Println("debe ser un numero")
		return 0
	}
	return operador1
}

// leer: lee lo que el usuario ingresa por teclado
// retorna la representacion en string
func leer() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// entry point
func main() {
	fmt.Println("indica la operacion")
	fmt.Println("cada elemento debe ir separado por un espacio ej: (2 + 5)")
	entrada := leer()
	c := calc{}
	c.operate(entrada)
}
