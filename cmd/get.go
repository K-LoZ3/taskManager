package cmd

import (
  "fmt"
  "Practicas/tareas/data"
  
  "github.com/spf13/cobra"
)

var getCmd = &cobra.Command {
  Use: "get",
  Short: "get tasks", //ver si esta bien escrito
  Long: "Muestra todas las tareas",
  Run: func(cmd *cobra.Command, arg []string) {
    //agregar tarea.
    err := data.InitDB()
    if err != nil {
      fmt.Println("Error al iniciar la base de datos", err)
    }
    defer data.Close()
    
    tareas, err := data.GetTask()
    if err != nil {
      fmt.Println(err)
      return
    }
    fmt.Println(tareas)
  },
}

func init() {
  
  rootCmd.AddCommand(getCmd)
}