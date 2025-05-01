package main

import "fmt"

type Mahasiswa struct {
	Nama string
	Kelas string
}

func main(){
	var data = map[string]Mahasiswa{
		"10148189": {
			"Ariel Rizki",
			"2KA04",
		},
		"10139190": {
			"Glenn Rhee",
			"9KA90",
		},
	}

	var search string
	fmt.Print("Masukkan NPM anda: ")
	fmt.Scanf("%s", &search)

	var value, oke = data[search]

	if oke {
		fmt.Printf("Nama anda %s \nKelas anda %s\n", value.Nama, value.Kelas)
	} else {
		fmt.Println("Upps, data tidak ditemukan!")
	}
}