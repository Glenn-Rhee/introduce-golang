package main
import "fmt"

func main(){
	for i := 1; i <= 13; i++ {
		if i % 2 != 0 {
			fmt.Printf("Angka %d\n", i)
		}
	}
}