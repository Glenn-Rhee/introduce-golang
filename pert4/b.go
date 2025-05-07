package main
import "fmt"

func main(){
	var buah_toko_ariel = [5]string{"Apel", "Semangka", "Jeruk", "Anggur", "Pepaya"}
	var toko_ariel = buah_toko_ariel[0:5]
	var buah_bulan_ini = [8]string{"Semangka", "Jeruk", "Anggur", "Rambutan", "Mangga", "Kelengkeng"}
	var katalog_bulanan = buah_bulan_ini[0:6]

	fmt.Println("Isi Toko Buah Ariel yaitu\t: ", toko_ariel)
	fmt.Println("Katalog buah bulan ini yaitu\t: ", katalog_bulanan)
	fmt.Println("Panjang Katalog Buah\t\t: ", len(toko_ariel))
	fmt.Println("Kapasitas Katalog Buah\t\t:", cap(buah_toko_ariel))
	fmt.Println("Panjang Katalog Buah Bulanan\t:", len(katalog_bulanan))
	fmt.Println("Kapasistas Katalog Buah Bulanan\t:", cap(katalog_bulanan))
}