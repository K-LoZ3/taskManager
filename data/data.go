package data

import (
  "fmt"
  "time"
  "database/sql"
  
  _ "modernc.org/sqlite"
)

var db *sql.DB

type Task struct {
  Id int
  TaskId int
  Name string
  Description string
  Check bool
  Date time.Time
}

func (t Task) String() string {
  status := "pendiente"
  if t.Check {
      status = "hecho"
  }

  return fmt.Sprintf(`
  (%d)-> %s [%s]
  %s
  Descripción: %s
  `,
      t.TaskId,
      t.Name,
      status,
      t.Date.Format("02-01-2006"), // formato de fecha YYYY-MM-DD
      t.Description,
  )
}

func InitDB() error {
  var err error
  
  db, err = sql.Open("sqlite", "task.db")
  if err != nil {
    return err
  }
  
  create := `
  CREATE TABLE IF NOT EXISTS task(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    taskid INTEGER NOT NULL,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    done BOOLEAN NOT NULL DEFAULT 0,
    date DATETIME DEFAULT CURRENT_TIMESTAMP
  );`
  
  _, err = db.Exec(create)
  if err != nil {
    return err
  }
  
  return nil
}

func Close() {
  db.Close()
}

func AddTask(t Task) error {
  _, err := db.Exec(`
    INSERT INTO task (taskid, name, description, date)
    VALUES (
      COALESCE((SELECT MAX(taskid) + 1 FROM task), 1),
      ?, ?, ?
  );`, t.Name, t.Description, t.Date)
    
  if err != nil {
    return err
  }
  
  return nil
}

func FindTaskId(id int) (Task, error) {
  var task Task
  
  err := db.QueryRow("SELECT name, description, done, date FROM task WHERE taskid = ?", id).Scan(&task.Name, &task.Description, &task.Check, &task.Date)
  
  task.TaskId = id
  
  if err != nil {
    return task, err
  }
  
  return task, err
}

func FindTaskName(name string) ([]Task, error) {
  var tasks []Task
  
  rows, err := db.Query("SELECT taskid, name, description, done, date FROM task WHERE name LIKE ?", "%"+name+"%")
  if err != nil {
    return tasks, err
  }
  defer rows.Close()
  
  for rows.Next() {
		var t Task
		var check int // SQLite guarda BOOL como INTEGER (0/1)

		err := rows.Scan(&t.TaskId, &t.Name, &t.Description, &check, &t.Date)
		if err != nil {
			return nil, err
		}

		t.Check = check == 1
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func GetTask() ([]Task, error) {
  var tasks []Task
  
  rows, err := db.Query("SELECT taskid, name, description, done, date FROM task")
  if err != nil {
    return tasks, err
  }
  defer rows.Close()
  
  for rows.Next() {
		var t Task
		var check int // SQLite guarda BOOL como INTEGER (0/1)

		err := rows.Scan(&t.TaskId, &t.Name, &t.Description, &check, &t.Date)
		if err != nil {
			return nil, err
		}

		t.Check = check == 1
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func CheckTask(id int) error {
  _, err := db.Exec("UPDATE task SET done = 1 - done WHERE taskid = ?", id)
    return err
}

func DeleteTask(id int) error {
  _, err := db.Exec("DELETE FROM task WHERE taskid = ?", id)
  if err != nil {
    return err
  }
  
  _, err = db.Exec(`
    WITH ordered AS (
      SELECT id, ROW_NUMBER() OVER (ORDER BY id) AS new_number
      FROM task
    )
    UPDATE task
    SET taskid = (
      SELECT new_number
      FROM ordered o
      WHERE o.id = task.id
    );
  `)
  if err != nil {
    fmt.Println(err)
  }
  
  return nil
}