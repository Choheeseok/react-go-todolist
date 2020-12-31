package model

// ToDo struct
type ToDo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Complete bool `json:"complete"`
}

// DBHandler interface
type DBHandler interface {
	GetToDos() []*ToDo
	AddToDo(toDo *ToDo) (*ToDo, error)
	DeleteToDo(id int) error
	CompleteToDo(id int) error
	GetDetail(id int) (*ToDo, error)
	Close()
}

// NewDBHandler return new DBHandler
func NewDBHandler() DBHandler {
	return newSqliteHandler()
}