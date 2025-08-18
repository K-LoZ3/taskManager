package cmd

import (
  "fmt"
  
  "github.com/spf13/cobra"
)

var addCmd = &cobra.Command {
  Use: "add",
  Short: "add task to list.", //ver si esta bien escrito
  Long: "Agrega una tarea a la lista de tareas",
  Run: func(cmd *cobra.Command, arg []string) {
    //agregar tarea.
    fmt.Println("add hola mundo")
  },
}

func init() {
  rootCmd.AddCommand(addCmd)
}