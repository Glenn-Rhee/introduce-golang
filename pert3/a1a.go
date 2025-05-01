package main 
import "fmt"

func main(){
	for i := 1; i <= 10; i ++ {
		if i % 2 == 0 {
			fmt.Printf("%d adalah bilangan genap\n", i)
			continue
		}

		fmt.Printf("%d adalah bilangan ganjil\n", i)
	}
}