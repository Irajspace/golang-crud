package sqlite

import (
	"database/sql"

	"github.com/irajspace/golang-crud/internal/config"
	_"github.com/mattn/go-sqlite3"
)


type SQLiteStorage struct {
	Db *sql.DB
}

func New(cfg *config.Config)(*SQLiteStorage, error){
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}
	_,err=db.Exec(`
	CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER NOT NULL,
		grade TEXT NOT NULL
	);
	`)
	if err != nil {
		return nil, err
	}
	return &SQLiteStorage{Db: db}, nil
}

func (s *SQLiteStorage) CreateStudent(name string, age int, grade string) (int64, error) {
	stmt,err:=s.Db.Prepare("INSERT INTO students(name,age,grade)VALUES(?,?,?)")
	if(err!=nil){
		return 0,err
	}
	defer stmt.Close()
	result,err:=stmt.Exec(name,age,grade)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id,nil	
	
}