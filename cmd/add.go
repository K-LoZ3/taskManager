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
    
    data.AddNameTask(arg[0])
    
    //Prueba para ver si almaceno bien
    tasks, err := data.GetTask()
    if err != nil {
      fmt.Println("Error al consultar las tareas", err)
    }
    
    tarea, err := data.FindTaskId(3)
    
    tarea3, err := data.FindTaskName("otra")
    
    fmt.Println(tasks, "tarea--->", tarea, "Busqueda--->", tarea3)
    //Pruebas de las funcione data
  },
}

func init() {
  rootCmd.AddCommand(addCmd)
}