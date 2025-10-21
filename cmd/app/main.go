package main

import (
	"fmt"

	"github.com/Nahid-ghorbani/graph-task-manager/initial/db"
)

func main() {

	database := db.Connect()
	_ = database

	
	fmt.Println("app works")
}