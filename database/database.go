package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var Database = func(db *gorm.DB) {
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
		return
	}

	dsname := os.Getenv("db_user")+":"os.Getenv("db_password")+"@tcp("+os.Getenv("db_server")
	+":"+os.Getenv("db_port")+")/"+os.Getenv("db_name")+"?charset+utf8mb4&parseTime=True&Loc=Local"

	if db, err := gorm.Open(mysql.Open(dsname), &gorm.Config{}); err != nil {
		fmt.Println("Error de Conexion a la BD")
		panic(err)
	} else {
		fmt.Println("Conexion a la BD")
		return db
	}
}()
