package main

import (
	"bufio"
	"os"
	"strconv"
)

func main2() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operacion := scanner.Text()

	valor1, error1 := strconv.Atoi(operacion)

	println(error1)
	println(operacion)
	println(valor1)
}
