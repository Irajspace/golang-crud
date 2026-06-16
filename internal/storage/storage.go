package storage


type Storage interface {
	CreateStudent(name string, age int, grade string) (int64, error)
}