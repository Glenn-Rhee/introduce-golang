package main

import "testing";

func TestFungsiHasil(t *testing.T){
	hasil := Hasil(4,2)
	if hasil != 8 {
		t.Errorf("Hasil perkalian 4 x 2 = 8, hasil yang di dapatkan %v", hasil)
	}
	
}