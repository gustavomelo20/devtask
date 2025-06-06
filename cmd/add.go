package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	tasks "devtask/internal/tasks"

	"github.com/spf13/cobra"
)

var (
	title       string
	description string
)

func AddTask(dbParam *sql.DB, title, description string) error {
	handler := tasks.NewAddTaskHandler(dbParam)

	cmd := tasks.AddTaskCommand{
		Title:       title,
		Description: description,
	}

	return handler.Handle(context.Background(), cmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adiciona uma nova tarefa",
	Run: func(cmd *cobra.Command, args []string) {
		if title == "" {
			fmt.Println("Você precisa informar um título com --title")
			os.Exit(1)
		}

		if err := AddTask(db, title, description); err != nil {
			fmt.Printf("Erro ao adicionar tarefa: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("Tarefa adicionada com sucesso!")
	},
}

func init() {
	addCmd.Flags().StringVarP(&title, "title", "t", "", "Título da tarefa (obrigatório)")
	addCmd.Flags().StringVarP(&description, "description", "d", "", "Descrição da tarefa")
	rootCmd.AddCommand(addCmd)
}
