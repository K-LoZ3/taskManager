package cmd

import (
  "fmt"
  "os"
  "bufio"
  "time"
  "strings"
  "Practicas/tareas/data"
  
  "github.com/spf13/cobra"
)

var custom bool
var descripcion string
var date string

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
    
    if len(arg) == 0 || arg[0] == "" {
      fmt.Println("Debe darle un nombre a la tarea")
      return
    }
    
    var newTask data.Task
    
    if date== "" {
      newTask = data.Task {
        Name: arg[0],
        Description: descripcion,
        Date: time.Now(),
      }
    
    } else {
      dateTask, err := time.Parse("02/01/2006", strings.TrimSpace(date))
      if err != nil {
        fmt.Println("error en la fecha", err)
      }
      newTask = data.Task {
        Name: arg[0],
        Description: descripcion,
        Date: dateTask,
      }
    }
    
    //newTask.Name = arg[0]
    
    if custom {
      reader := bufio.NewReader(os.Stdin)
      fmt.Print("Descripcion: ")
      newTask.Description, _ = reader.ReadString('\n')
      fmt.Print("Fecha para la tarea dd/mm/yyyy: ")
      dateStr, _ := reader.ReadString('\n')
      
      if dateStr == "" {
        newTask.Date = time.Now()
      } else {
        newTask.Date, err = time.Parse("02/01/2006", strings.TrimSpace(dateStr))
      }
      
      if err != nil {
        fmt.Println("Error en la fecha", err)
      }
      // terminar la logica para guardar los datos
    }
    
    if newTask.Date.IsZero() {
      newTask.Date = time.Now()
    }
    
    data.AddTask(newTask)
    
    //para probar como va guardando mientras
    //tareas, err := data.GetTask()
    //fmt.Println(tareas, err)
  },
}

func init() {
  
  addCmd.Flags().StringVarP(&descripcion, "descrip", "d", "", "Ingresa la descripcion")
  
  addCmd.Flags().StringVarP(&date, "fecha", "f", "", "Ingresa la fecha")
  
  addCmd.Flags().BoolVarP(&custom, "custom", "c", false, "Marca para saber si se usa el metodo completoo no.")
  
  rootCmd.AddCommand(addCmd)
}