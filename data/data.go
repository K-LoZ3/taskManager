package data

import (
  "time"
  "database/sql"
  
  _ "modernc.org/sqlite"
)

var db *sql.db

type Task struct {
  Id int
  Name string
  Description string
  Check bool
  Date time.Time
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
  name TEXT NOT NULL,
  description TEXT NOT NULL,
  check BOOLEAN NOT NULL DEFAULT 0,
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

func AddNameTask(t string) error {
  _, err = db.Exec("INSERT INTO task(name, description) VALUES( ?, ?)", t, "")
  if err != nil {
    return err
  }
  
  return nil
}

func AddTask(name string, description string, date time.Time) error {
  _, err = db.Exec("INSERT INTO task(name, description, date) VALUES( ?, ?, ?)", name, description, date)
  if err != nil {
    return err
  }
  
  return nil
}

func FindTaskId(id int) (Task, error) {
  var task Task
  
  err := db.QueryRow("SELECT name, description, check, date FROM task WHERE id = ?", id).Scan(&task.Name, &task.Description, &task.Check, &task.Date)
  if err != nil {
    return task, err
  }
  
  return task, err
}

func FindTaskName(name) (Task, error) {
  var task Task
  
  err := db.QueryRow("SELECT id, description, check, date FROM task WHERE name = ?", name).Scan(&task.Id, &task.Description, &task.Check, &task.Date)
  if err != nil {
    return task, err
  }
  
  return task, err
}

func Task() ([]Task, error) {
  var tasks []Task
  
  rows, err := db.Query("SELECT id, description, check, date FROM task")
  if err != nil {
    return task, err
  }
  defer rows.Close()
  
  for rows.Next() {
		var t Task
		var check int // SQLite guarda BOOL como INTEGER (0/1)

		err := rows.Scan(&t.Id, &t.Name, &t.Description, &check, &t.Date)
		if err != nil {
			return nil, err
		}

		t.Check = Check == 1
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}