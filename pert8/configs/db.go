package configs

import (
	"fmt"
	"log"
	"time"

	// TODO: ACT NO 4 IMPORT PACKAGE
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	user := "cpc18"
	password := "lepkom@123"
	host := "dbms.lepkom.f4.com"
	dbname := "event_realm"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, host, dbname)
	log.Printf("Connecting to database: %s:%s@tcp(%s)/%s", user, password, host, dbname)
	// Tambahkan parameter parseTime=true di string koneksi
	// untuk menangani konversi waktu MySQL ke Go time.Time
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal membuka koneksi database:", err)
	}

	// Konfigurasi pool koneksi
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(time.Minute * 5)

	// Verifikasi koneksi dengan ping
	err = DB.Ping()
	if err != nil {
		log.Fatal("Gagal melakukan ping ke database:", err)
	}

	// Verifikasi tabel events
	_, err = DB.Query("SELECT 1 FROM events LIMIT 1")
	if err != nil {
		log.Fatal("Gagal menjalankan query pada tabel events:", err)
	}

	log.Println("Berhasil terhubung ke database")
}
