package cmd

import (
  "fmt"
  "strconv"
  "Practicas/tareas/data"
  
  "github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command {
  Use: "delete",
  Short: "delete task to list.", //ver si esta bien escrito
  Long: "Elimina una tarea a la lista de tareas",
  Run: func(cmd *cobra.Command, arg []string) {
    //agregar tarea.
    err := data.InitDB()
    if err != nil {
      fmt.Println("Error al iniciar la base de datos", err)
    }
    defer data.Close()
    
    if len(arg) == 0 || arg[0] == "" {
      fmt.Println("Ingresar un id para borrar, si no sabe usa get para ver la lista.")
      return
    }
    
    id, err := strconv.Atoi(arg[0])
    if err != nil {
      fmt.Println("El id debe ser un numero entero.", err)
      return
    }
    
    data.DeleteTask(id)
    
    //para probar como va guardando mientras
    //tareas, err := data.GetTask()
    //fmt.Println(tareas, err)
  },
}

func init() {
  
  rootCmd.AddCommand(deleteCmd)
}