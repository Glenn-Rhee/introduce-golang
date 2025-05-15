package main

import (
	"fmt"
	"strings"
)

func add(a int, b int) int{
	return a + b
}

func repeat(str string, count int) string {
	return strings.Repeat(str, count)
}

func getDimensions() (int, int){
	return 100, 50
}

func main(){
	a := 10
	b := 5
	hasil := add(a, b)
	fmt.Printf("Hasil dari %d + %d = %d\n", a, b, hasil)

	repeated := repeat("Ariel", 10)
	fmt.Println("String repeated: ", repeated)

	width, _ := getDimensions()
	fmt.Println("Dimension: ", width)
}
