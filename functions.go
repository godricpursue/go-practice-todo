package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
}

func getDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		GetEnv("DB_HOST"), GetEnv("DB_USER"), GetEnv("DB_PASSWORD"), GetEnv("DB_NAME"), GetEnv("DB_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func GetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic(fmt.Sprintf("%s is not set in the .env file", key))
	}
	return value
}

func addTodo(db *gorm.DB, desc string) {
	if tx := db.Create(&TodoInput{Input: desc, IsDone: false}); tx.Error != nil {
		panic("Todo not added")
	}
}

func deleteTodo(db *gorm.DB, id uint) {
	if tx := db.Delete(&TodoInput{ID: id}); tx.Error != nil {
		panic("Todo not deleted")
	}
}

func toggleIsDone(db *gorm.DB, id uint) {
	var todo TodoInput
	if tx := db.First(&todo, id); tx.Error != nil {
		panic("Todo not found")
	}

	// Toggle the IsDone field
	todo.IsDone = !todo.IsDone
	if tx := db.Save(&todo); tx.Error != nil {
		panic("Failed to update IsDone")
	}
}
