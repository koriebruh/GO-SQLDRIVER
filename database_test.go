package go_mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //<-- menginisialisasi driver mysql
	"testing"
)

// <-- try make connection
func TestConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root:korie123@tcp(localhost:3306)/belajar_golang_db")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close() // <-- menutuo golangdb
	fmt.Println("success connec")
}

func TestFuncConnection(t *testing.T) {
	db := GetConnection()
	defer db.Close()
	fmt.Println("Success connect")
}
