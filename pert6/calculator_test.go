package main

import (
	"fmt"
	"testing"
)

func TestPenjumlahan(t *testing.T){
	a := 5
	b := 3
	hasil := penjumlahan(a, b)
	expected := a + b

	if hasil != expected {
		t.Errorf("%d + %d should be %d. But got %d", a, b, expected, hasil)
	}

	fmt.Println("Test fungsi Penjumlahan success")
}

func TestPengurangan(t *testing.T){
	a := 5
	b := 3
	hasil := pengurangan(a, b)
	expected := a - b

	if hasil != expected {
		t.Errorf("%d - %d should be %d. But got %d", a, b, expected, hasil)
	}

	fmt.Println("Test fungsi pengurangan success")

}

func TestPerkalian(t *testing.T){
	a := 5
	b := 3
	hasil := perkalian(a, b)
	expected := a * b

	if hasil != expected {
		t.Errorf("%d * %d should be %d. But got %d", a, b, expected, hasil)
	}
	
	fmt.Println("Test fungsi perkalian success")

}