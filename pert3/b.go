package main
import "fmt"

func main(){
	var gaji int
	var grade string
	var totalGaji float64

	fmt.Print("Masukkan gaji\t: ")
	fmt.Scan(&gaji)

	if gaji <= 2000000 {
		grade = "A"
		totalGaji = float64(gaji) + (float64(gaji) * 0.2)
	} else if gaji > 2000000 && gaji <= 3000000 {
		grade = "B"
		totalGaji = float64(gaji) + (float64(gaji) * 0.15)
	} else if gaji > 3000000 && gaji <= 5000000 {
		grade = "C"
		totalGaji = float64(gaji) + (float64(gaji) * 0.1)		
	} else {
		grade = "D"
		totalGaji = float64(gaji) + (float64(gaji) * 0.05)
	}

	fmt.Println("")

	fmt.Printf("Gaji awal\t: Rp %d\n", gaji)
	fmt.Printf("Grade Tunjangan\t: %s\n", grade)
	fmt.Printf("Total Gaji\t: Rp %.2f\n", totalGaji)
}