package cmd

import (
	"context"
	"fmt"
	"log"

	tasks "devtask/internal/tasks"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todas as tarefas",
	Run: func(cmd *cobra.Command, args []string) {
		handler := tasks.NewListTasksHandler(db)

		taskList, err := handler.Handle(context.Background(), tasks.ListTasksCommand{})
		if err != nil {
			log.Fatalf("Erro ao listar tarefas: %v", err)
		}

		fmt.Println("Lista de tarefas:")
		for _, t := range taskList {
			status := "Pendente"
			if t.Done {
				status = "Concluída"
			}
			fmt.Printf("ID: %d\nTítulo: %s\nDescrição: %s\nStatus: %s\nCriada em: %s\n\n",
				t.ID, t.Title, t.Description, status, t.CreatedAt)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
