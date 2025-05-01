package main

import "fmt"

var nilai1, nilai2 int

func main(){
	defer fmt.Println("-----SELESAI-----")
	fmt.Print("Masukkan bilangan 1: ")
	fmt.Scan(&nilai1)
	fmt.Print("Masukkan bilangan 2: ")
	fmt.Scan(&nilai2)

	hasil := nilai1 % nilai2

	fmt.Printf("Hasil dari sisa pembagian nilai 1 dengan nilai 2 = %d \n", hasil)
}