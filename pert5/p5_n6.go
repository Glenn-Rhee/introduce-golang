package main
import "fmt"

func processFile(filename string){
	fmt.Println("Opening file:", filename)
	defer fmt.Println("Closing file:", filename)
	fmt.Println("Processing file...")
}

func executeTransaction(userId string, isSuccess bool) {
	fmt.Println("Starting transaction for user:", userId)
	defer fmt.Println("Transaction finished")

	fmt.Println("Connected to database")
	defer fmt.Println("Disconnected from database")

	fmt.Println("Beginning transaction")
	if !isSuccess {
		defer fmt.Println("Rolling back transaction")
	}

	fmt.Println("Precessing transaction...")
	
	if isSuccess {
		fmt.Println("Committing transaction")
		defer fmt.Println("Transaction comitted")
	} else {
		fmt.Println("Transaction failed")
		return
	}
}

func main(){
	processFile("nama_npm.txt")

	fmt.Println("\nSucessful transction:")
	executeTransaction("user178", true)

	fmt.Println("\nFailed transaction:")
	executeTransaction("user178", false)
}