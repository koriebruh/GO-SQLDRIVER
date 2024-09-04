package go_mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestOperationSQL(t *testing.T) { // <-- insert,update,delete (yang tidak mengembalikan hasil)
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	// <-- better gunaan tanda tanya unutk isi value
	script := "INSERT INTO customer(id, name) VALUES ('nathalie','Nathalie Yovela');"

	_, err := db.ExecContext(ctx, script)
	if err != nil {
		panic(err)
	}
	fmt.Println("Congrats success insert new customer")
}

// <-- alter alter table
func TestQuerySQL(t *testing.T) { // <-- mengembalikan result
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	query := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		panic(err)
	}

	// <-- ngambil data
	for rows.Next() {
		var id, name string
		var email sql.NullString //<-- tipe data custom unutk mengangani data null
		var balance int32
		var rating float64 // <--same double
		var birthDate sql.NullTime
		var createdAt time.Time
		var married bool
		//"SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}
		fmt.Println("Id :", id)
		fmt.Println("Name :", name)
		if email.Valid { // <-- mengambik data jika tipedata sql.Mullxxx
			fmt.Println("Email :", email.String)
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating :", rating)
		if birthDate.Valid {
			fmt.Println("Birth_date :", birthDate.Time)
		}
		fmt.Println("Created_at :", createdAt)
		fmt.Println("Married :", married)
		fmt.Println("")

	}

	defer rows.Close()

}

func TestInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	username := "admin"
	password := "admin"

	// <-- test sql injection
	//username := "admin'; #"
	//password := "wrong"

	ctx := context.Background()

	sqlQuery := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"
	rows, err := db.QueryContext(ctx, sqlQuery, username, password)
	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			panic(err)
		} else {
			fmt.Println("Success login", username)
		}

	} else {
		fmt.Println("Error Login")
	}

}

// <-- last insert id autoincrement
func TestAutoInc(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	email := "bocil dino"
	comment := "bang kapan dinoo"

	ctx := context.Background()
	sqlQuery := "INSERT INTO comments(email,comment) VALUES (?,?)"
	result, err := db.ExecContext(ctx, sqlQuery, email, comment)
	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	fmt.Println("Success Insert new comment with Id := ", insertId)
}

/*
SAAT KITA MENGUNAKAN ExecContext dan QueryContext tidak efisien karena menanyakan koneksi sebelum mekaukan query
padahal format query sama maka dari itu ada PrepareContext yang memastikan 1 koneksi
*/

func TestPrepareContext(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	sqlQuery := "INSERT INTO comments (email,comment) VALUES (?,?)"
	statement, err := db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "test" + strconv.Itoa(i) + "gmail.com"
		comment := "gg kata ilham " + strconv.Itoa(i)
		result, err := statement.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("Last comment Id :=", lastInsertId)
	}

}

/*
flow TRANSACTION SQL :
1. start transaction, memulai transaksi
2. saat kita melakukan transacsi kan tidak hanya ada 1 tabel yg dilakukan perubhanan (Query)
3. jika salah 1 query error kita bisa melakuakn rollback (query setelah start transaction akan di batalkan)
4. jika tidak ada error kita bisa commit untuk menyimpan perubahan transaksi
*/

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	// <-- start transaction
	sqlQuery := "INSERT INTO comments (email,comment) VALUES (?,?)"
	for i := 0; i < 10; i++ {
		email := "test" + strconv.Itoa(i) + "gmail.com"
		comment := "gg kata ilham " + strconv.Itoa(i)
		result, err := tx.ExecContext(ctx, sqlQuery, email, comment)

		if err != nil {
			panic(err)
		}
		lastInsertId, _ := result.LastInsertId()
		fmt.Println("Last comment Id :=", lastInsertId)
	}

	if err := tx.Commit(); err != nil { // <-- ketika sukses akan commit(disimpan)
		tx.Rollback() // <-- ketika proses query error akan melakukan rollback(semua transaksi sebelumnya akan dibatalkan)
	}

}
