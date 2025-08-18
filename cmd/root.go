package cmd

import (
  "fmt"
  
  "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
  Use: "task",
  Short: "use task with cobra-cli", //ver si esta bien escrito
  Long: "Task es para administrar una lista de tareas",
  Run: func(cmd *cobra.Command, arg []string) {
    //agregar tarea.
    fmt.Println("task hola mundo")
  },
}

func Execute() {
  rootCmd.Execute()
}
