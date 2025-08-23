package cmd

import (
  "fmt"
  "strconv"
  "Practicas/tareas/data"
  
  "github.com/spf13/cobra"
)

var updateCmd = &cobra.Command {
  Use: "update",
  Short: "update 'done' task to list.", //ver si esta bien escrito
  Long: "Actualiza el checkeo de una tarea",
  Run: func(cmd *cobra.Command, arg []string) {
    //agregar tarea.
    err := data.InitDB()
    if err != nil {
      fmt.Println("Error al iniciar la base de datos", err)
    }
    defer data.Close()
    
    if len(arg) == 0 || arg[0] == "" {
      fmt.Println("Ingresar un id para actualizar, si no sabe usa get para ver la lista.")
      return
    }
    
    id, err := strconv.Atoi(arg[0])
    if err != nil {
      fmt.Println("El id debe ser un numero entero.", err)
      return
    }
    
    data.CheckTask(id)
    
    //para probar como va guardando mientras
    //tareas, err := data.GetTask()
    //fmt.Println(tareas, err)
  },
}

func init() {
  
  rootCmd.AddCommand(updateCmd)
}