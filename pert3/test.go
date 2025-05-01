package main

import "fmt"

func main(){
	var grade  = "A"

	switch grade {
	case "A":
		fmt.Println("Nilai kamu bagus banget")
		break;
	case "B": 
		fmt.Println("Nilai kamu lumayan")
		break;
	default:
		fmt.Println("Kerja keras lagi ya")
		break;
	}
}