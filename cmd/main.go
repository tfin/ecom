package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/tfin/ecom/cmd/api"
	"github.com/tfin/ecom/configs"
	"github.com/tfin/ecom/db"
)

func main() {

	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 configs.Envs.DBUser,
		Passwd:               configs.Envs.DBPassword,
		Addr:                 configs.Envs.DBAddress,
		DBName:               configs.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
