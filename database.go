package go_mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //<-- menginisialisasi driver mysql
	"time"
)

// <-- create func connection

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:korie123@tcp(127.0.0.1:3306)/belajar_golang_db?parseTime=true")
	// <-- ?parseTime=true , di tambahkan itu jika nanti parse ke golang otomatis jadi bisa tipe data Time.Time
	if err != nil {
		panic(err.Error())
	}

	db.SetMaxIdleConns(10)                  //<-- berapa minimal koneksi saat pertama jalankan appnya
	db.SetMaxOpenConns(100)                 //<-- maxsimal koneksi,
	db.SetConnMaxIdleTime(5 * time.Minute)  //<-- berapa lama koneksi yang sudah tidak digunakan akan dihapus
	db.SetConnMaxLifetime(60 * time.Minute) //<-- seberapa lama koneksi boleh digunakan

	return db
}
