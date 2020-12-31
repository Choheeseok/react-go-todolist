package model

import (
	"database/sql"
	"errors"
	"fmt"

	// to use sqlite3
	_ "github.com/mattn/go-sqlite3"
)


type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) GetToDos() []*ToDo {
	toDos := []*ToDo{}
	rows, err := s.db.Query("SELECT id, text, complete FROM todos")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var toDo ToDo
		rows.Scan(&toDo.ID, &toDo.Text, &toDo.Complete)
		fmt.Println(toDo.ID)
		toDos = append(toDos, &toDo)
	}
	return toDos
}

func (s *sqliteHandler) AddToDo(toDo *ToDo) (*ToDo, error) {
	stmt, err := s.db.Prepare("INSERT INTO todos (id, text, complete) VALUES (?, ?, ?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(toDo.ID, toDo.Text, toDo.Complete)
	if err != nil {
		panic(err)
	}
	return toDo, nil
} 

func (s *sqliteHandler) DeleteToDo(id int) error {
	stmt, err := s.db.Prepare("DELETE FROM todos WHERE id=?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	rst, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	cnt, _ := rst.RowsAffected()
	if cnt > 0 {
		return nil
	} 
	return errors.New("ID does not exist")
}

func (s *sqliteHandler) CompleteToDo(id int) error{
	stmt, err := s.db.Prepare("UPDATE todos SET completed=? WHERE id=?")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	rst, err := stmt.Exec(false, id)
	if err != nil {
		panic(err)
	}
	cnt, _ := rst.RowsAffected()
	if cnt > 0 {
		return nil
	}
	return errors.New("ID does not exist")
}

func (s *sqliteHandler) GetDetail(id int) (*ToDo, error) {
	row := s.db.QueryRow("SELECT id, text, complete complete FROM todos WHERE id=?",id)
	var toDo ToDo
	err := row.Scan(&toDo.ID, &toDo.Text, &toDo.Complete)
	if err != nil {
		panic(err)
	}
	fmt.Println(toDo)
	
	return &toDo,nil
}

func (s *sqliteHandler) Close() {
	s.db.Close()
}

func newSqliteHandler() DBHandler {
	database, err := sql.Open("sqlite3", "./test.db" )
	if err != nil{
		panic(err)
	}

	statement, _ := database.Prepare(
		`CREATE TABLE IF NOT EXISTS todos (
			id	INTEGER PRIMARY KEY,
			text	TEXT,
			complete BOOLEAN)`)
	statement.Exec()
	statement.Close()

	return &sqliteHandler{db: database}
}