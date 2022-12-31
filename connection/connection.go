package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func DatabaseConnect() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	databaseUrl := "postgres://postgres:root@localhost:5432/dbproject"

	var err error
	Conn, err = pgx.Connect(context.Background(), databaseUrl)

	if err != nil {
		//Stderr untuk mengubah ke standar inputan
		fmt.Fprintf(os.Stderr, "Koneksi database gagal: %v\n", err)
		os.Exit(1) //jika tidak berjalan maka akan menutup
	}

	fmt.Println("koneksi database berhasil !!")

}
