package main
import "fmt"

func operate(a int, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func main(){
	
	result := operate(5, 3, func(a, b int) int {
		return a + b
	})

	fmt.Println("5 + 3 =", result)

	result = operate(5, 3, func(a, b int) int {
		return a * b
	})
	fmt.Println("5 * 3 =", result)

	result = operate(5, 3, func (a, b int) int {
		return 5 - 3
	})
	fmt.Println("5 - 3 =", result)
}