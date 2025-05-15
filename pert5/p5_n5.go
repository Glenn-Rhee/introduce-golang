package main
import "fmt"

type Counter struct {
	count int
}

func (c Counter) Increment(){
	c.count++
}

func (c *Counter) IncrementByPointer(){
	c.count++
}

func (c Counter) GetValue() int{
	return c.count
}

func main(){
	counter := Counter{count: 10}
	counter.Increment()
	fmt.Println("After increment:", counter.GetValue())

	counter.IncrementByPointer()
	fmt.Println("After increment by pointer:", counter.GetValue())

	c := &Counter{count: 10}
	c.Increment()
	fmt.Println("Pointer after increment:", c.GetValue())
	
	c.Increment()
	fmt.Println("Pointer after Increment by Pointer:", c.GetValue())
}