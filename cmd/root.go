package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "devtask",
	Short: "Um gerenciador de tarefas simples",
	Long:  "Aplicação CLI para gerenciar tarefas com SQLite.",
}

var db *sql.DB

func SetDB(database *sql.DB) {
	db = database
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
