package main

import (
	"fmt"
	"gorm.io/gorm"
)

type TodoInput struct {
	ID     uint `gorm:"primarykey"`
	Input  string
	IsDone bool
}

var db *gorm.DB

func main() {
	loadEnv()

	var err error
	db, err = getDB()
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("db connection success")

	if err := db.AutoMigrate(&TodoInput{}); err != nil {
		panic("failed to migrate")
	}

	//desc := "query"
	//addTodo(db, desc)
	//deleteTodo(db, 8)
	//toggleIsDone(db, 9)

	var todolist []TodoInput
	tx := db.Find(&todolist)
	if tx.Error != nil {
		panic("failed to query")
	}

	for _, todo := range todolist {
		fmt.Println(todo)
	}
}
