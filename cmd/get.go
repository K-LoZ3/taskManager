package cmd

import (
  "fmt"
  "Practicas/taskManager/data"
  
  "github.com/spf13/cobra"
)

type tasks []data.Task

func (ts tasks) String() string {
  var out string
  for _, t := range ts {
    out += t.String()
  }
  return out
}

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
    
    var tareas tasks
    
    tareas, err = data.GetTask()
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