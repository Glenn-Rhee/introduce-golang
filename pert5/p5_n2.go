package main

import (
	"errors"
	"fmt"
)

func calculateStats(numbers ...int) (min, max int, avg float64) {
	if len(numbers) == 0 {
		return 0, 0, 0
	}
	min = numbers[0]
	max = numbers[0]
	sum := 0

	for _, num := range numbers {
		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
		sum += num
	}

	avg = float64(sum) / float64(len(numbers))
	return 
}

func divide(dividend, divisor int) (float64, error){
	if divisor == 0 {
		return 0, errors.New("Tidak bisa dibagi dengan nol")
	}

	return float64(dividend) / float64(divisor), nil
}

func main(){
	min, max, avg := calculateStats(4, 17, 9, 12, 6)
	fmt.Printf("Stats - min: %d, max: %d, max: %.2f\n", min, max, avg)

	divided, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", divided)
	}

	divided, err = divide(14, 0)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("14 /  = %.2f\n", divided)
	}
}