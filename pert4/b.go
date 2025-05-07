package main
import "fmt"

func main(){
	tokoAriel := []string{"Apel", "Semangka", "Jeruk", "Anggur", "Pepaya"}
	fmt.Println("Isi Toko Buah Ariel yaitu: ", tokoAriel)

	var katalogBulanan = tokoAriel[1:4]
	katalogBulanan = append(katalogBulanan, "Rambutan", "Mangga", "Kelengkeng")
	fmt.Println("Katalog buah bulan ini yaitu: ", katalogBulanan)

	fmt.Println("Panjang katalog Buah: ", len(tokoAriel))
	fmt.Println("Kapasitas katalog buah: ", cap(tokoAriel))
	fmt.Println("Panjang katalog buah bulanan: ", len(katalogBulanan))
	fmt.Println("kapasitas katalog buah bulanan: ", cap(katalogBulanan))
	
}