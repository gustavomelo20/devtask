package main

import (
	"devtask/cmd"
	"devtask/internal/db"
)

func main() {
	db.InitDB()
	cmd.SetDB(db.DB)
	cmd.Execute()
}
