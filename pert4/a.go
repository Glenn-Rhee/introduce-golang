package main
import "fmt"

func main(){
	var gajiSekarang, ekspektasiGaji int

	fmt.Print("Masukkan gaji anda: ")
	fmt.Scan(&gajiSekarang)

	fmt.Print("Masukkan gaji yang anda inginkan: ")
	fmt.Scan(&ekspektasiGaji)

	gajiSekarang = *&ekspektasiGaji

	fmt.Print("Gaji anda sekarang ", gajiSekarang)
}