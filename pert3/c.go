package main 
import "fmt"

func main(){
	var i, j int

	for j = 0; j < 10; j++ {
		fmt.Print("Masukkan angka: ")
		fmt.Scan(&i)

		if i > 10{
			fmt.Println("Pertanyaan selesai karena nilai I lebih besar dari 10")
			break
		}
		
		if i % 2 == 0 {
			fmt.Printf("%d adalah Genap\n", i)
			continue
		}

		fmt.Printf("%d adalah Ganjil\n", i)
	}
}