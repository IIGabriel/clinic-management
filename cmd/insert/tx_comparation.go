package main

import (
	"fmt"
	"github.com/IIGabriel/clinic-management/internal/models"
	"github.com/IIGabriel/clinic-management/pkg/db/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {

	db := mysql.Mysql()

	fmt.Println("Inserindo registros sem transação...")
	startNoTx := time.Now()
	insertWithoutTransaction(db)
	durationNoTx := time.Since(startNoTx).Seconds()
	fmt.Printf("Tempo sem transação: %.2f segundos\n", durationNoTx)

	fmt.Println("Inserindo registros com transação...")
	startTx := time.Now()
	insertWithTransaction(db)
	durationTx := time.Since(startTx).Seconds()
	fmt.Printf("Tempo com transação: %.2f segundos\n", durationTx)

}

func insertWithoutTransaction(db *gorm.DB) {
	for i := 0; i < 10000; i++ {
		patient := models.Patient{
			FirstName: fmt.Sprintf("John %d", i),
			LastName:  fmt.Sprintf("Doe %d", i),
			BirthDate: time.Now().AddDate(0, 0, i),
			Address:   fmt.Sprintf("123 Main St %d", i),
			Phone:     fmt.Sprintf("555-5555 %d", i),
			Email:     fmt.Sprintf("teste%d@gmail.com", i),
		}
		db.Create(&patient)
	}
}

func insertWithTransaction(db *gorm.DB) {
	tx := db.Begin()
	if tx.Error != nil {
		log.Fatalf("Erro ao iniciar a transação: %v", tx.Error)
	}

	for i := 0; i < 10000; i++ {
		patient := models.Patient{
			FirstName: fmt.Sprintf("John %d 2", i),
			LastName:  fmt.Sprintf("Doe %d 2", i),
			BirthDate: time.Now().AddDate(0, 0, i),
			Address:   fmt.Sprintf("123 Main St %d 2", i),
			Phone:     fmt.Sprintf("555-5555 %d 2", i),
			Email:     fmt.Sprintf("teste%d-2@gmail.com", i),
		}
		if err := tx.Create(&patient).Error; err != nil {
			tx.Rollback()
			log.Fatalf("Erro ao inserir registro: %v", err)
		}
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		log.Fatalf("Erro ao cometer a transação: %v", err)
	}
}
