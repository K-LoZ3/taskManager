package cmd

import (
  "fmt"
  "Practicas/tareas/data"
  
  "github.com/spf13/cobra"
)

var addCmd = &cobra.Command {
  Use: "add",
  Short: "add task to list.", //ver si esta bien escrito
  Long: "Agrega una tarea a la lista de tareas",
  Run: func(cmd *cobra.Command, arg []string) {
    //agregar tarea.
    err := data.InitDB()
    if err != nil {
      fmt.Println("Error al iniciar la base de datos", err)
    }
    defer data.Close()
    
    data.AddNameTask("Nueva tareas de prueba hardcodeada")
    
    //Prueba para ver si almaceno bien
    tasks, err := data.GetTask()
    if err != nil {
      fmt.Println("Error al consultar las tareas", err)
    }
    
    fmt.Println(tasks)
  },
}

func init() {
  rootCmd.AddCommand(addCmd)
}