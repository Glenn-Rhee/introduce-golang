package main

import (
	"fmt"
	"strconv"
)

func main(){
	var nama_depan = "Ariel"
	var nama_belakang = "Rizki"
	var umur = 20
	var bil = strconv.Itoa(umur)

	fmt.Print("My name is " + nama_depan + " " + nama_belakang + ". Umur saya " + bil)
}