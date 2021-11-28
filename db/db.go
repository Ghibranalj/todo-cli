package db

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Ghibranalj/todo-cli/utils"
)

type database struct {
	Total int      `json:"Total"`
	Todos []string `json:"Todos"`
}

var (
	db     database
	dbPath string
)

func Init(path string) error {
	dbPath = path + "/todos.json"

	emptydb, err := json.Marshal(database{0, []string{}})
	if err != nil {
		return err
	}
	err = utils.CreateFile(dbPath, string(emptydb))
	if err != nil {
		return err
	}

	dbstr, err := utils.ReadFile(dbPath)

	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(dbstr), &db)

	if err != nil {
		return err
	}
	return nil
}

func Save() error {
	db.Total = len(db.Todos)
	dbstr, err := json.Marshal(db)

	if err != nil {
		return err
	}

	return utils.WriteToFile(dbPath, string(dbstr))
}

func Read(i int) string {
	return db.Todos[i]
}

func ReadAll() []string {
	return db.Todos
}

func Add(todo string) {
	db.Todos = append(db.Todos, todo)
}

func Replace(i int, content string) error {

	db.Todos[i] = content
	return nil
}

func Del(i int) error {
	if i >= db.Total {
		errstr := fmt.Sprintf("thre is only %d todos", db.Total)
		return errors.New(errstr)
	}
	// Remove the element at index i
	copy(db.Todos[i:], db.Todos[i+1:])    // Shift a[i+1:] left one index.
	db.Todos[len(db.Todos)-1] = ""        // Erase last element (write zero value).
	db.Todos = db.Todos[:len(db.Todos)-1] // Truncate slice.
	return nil
}

func Size() int {
	return len(db.Todos)
}
