package connection

import (
	"github.com/jinzhu/gorm"

	//Iniciando postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func AbrirConexao() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=gotuto sslmode=disable password=vili")
	if err != nil {
		println(err)
		panic("Falha ao conectar ao banco de dados")
	}
	return db, err
}
