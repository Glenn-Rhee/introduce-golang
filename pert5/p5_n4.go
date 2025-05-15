package main
import "fmt"

type Rectangle struct {
	width int
	height int
}

func (r Rectangle) Area() int {
	return r.width * r.height
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.width + r.height)
}

func (r *Rectangle) Scale(s int) {
	r.width *= s
	r.height *= s
}

func main(){
	persegiPanjang := Rectangle{width: 5, height: 3}
	fmt.Println("Persegi panjang:", persegiPanjang)

	area := persegiPanjang.Area()
	fmt.Println("Luas:", area)

	perimeter := persegiPanjang.Perimeter()
	fmt.Println("Perimeter:", perimeter)

	persegiPanjang.Scale(2)
	fmt.Println("After scaling:", persegiPanjang)
	fmt.Println("New area:", persegiPanjang.Area())
}