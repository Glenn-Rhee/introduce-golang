package main
import (
	"fmt"
	"testing"
)

func TestPembagian(t *testing.T){
	a := 10
	b := 2
	expected := float64(a) / float64(b)

	hasil := pembagian(a, b)
	if hasil != expected {
		t.Errorf("%d / %d should be %.2f. But got %.2f", a, b, expected, hasil)
	}

	fmt.Println("Test fungsi pembagian success")
}