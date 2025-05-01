package main 

import "fmt"

const PI = 22/7

var jari int

func main(){
	fmt.Print("Masukkan jari-jari lingkaran = ")
	fmt.Scan(&jari)
	luas := PI * (jari * jari)
	fmt.Println("Luas lingkaran =", luas)
}