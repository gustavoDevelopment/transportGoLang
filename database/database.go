package database

import (
	"fmt"
	"os"
	"restMux/logger"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database = func() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("Error al cargar el archivo .env: %v", err))
	}

	dsname := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("db_user"),
		os.Getenv("db_password"),
		os.Getenv("db_server"),
		os.Getenv("db_port"),
		os.Getenv("db_name"),
	)

	db, err := gorm.Open(mysql.Open(dsname), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Error de conexión a la BD: %v", err))
	}

	//err = db.AutoMigrate(&catalog.EPS{})
	//err = db.AutoMigrate(&catalog.Department{})
	//err = db.AutoMigrate(&catalog.DocumentType{})
	//err = db.AutoMigrate(&catalog.City{})
	//err = db.AutoMigrate(&catalog.InsuranceCompany{})
	//err = db.AutoMigrate(&catalog.TruckBrand{})
	//err = db.AutoMigrate(&model.Person{})
	//err = db.AutoMigrate(&model.Truck{})
	//if err != nil {
	//		log.Fatalf("Error en la migración: %v", err)
	//	}

	logger.Log.Info("Conexión a la BD exitosa")
	return db
}()
