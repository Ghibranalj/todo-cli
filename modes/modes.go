package modes

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Ghibranalj/todo-cli/db"
	"github.com/Ghibranalj/todo-cli/utils"
	"github.com/charmbracelet/glamour"
)

var filePath, editor string

const (
	// 	head = ` _______   ____    _____     ____
	// |__  __|  / __ \  |  __ \   / __ \
	//   | |    | |  | | | |  | | | |  | |
	//   | |    | |  | | | |  | | | |  | |
	//   | |    | |__| | | |__| | | |__| |
	//   |_|     \____/  |_____/   \____/ `
	colorOrange = "\033[38;5;166m"
	colorPurple = "\033[35m"
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	head        = `*****  ****  ****   ****
   *    *  *  *   *  *  *
   *    *  *  *   *  *  *
   *    ****  ****   ****`
)

func check(e error) {
	if e != nil {
		fmt.Printf("Error: %s \n", e.Error())
		os.Exit(1)
	}
}

func Print(number string) {
	i, err := strconv.Atoi(number)
	i--
	check(err)
	todo := db.Read(i)

	out, err := glamour.Render(todo, "dark")
	check(err)

	fmt.Println(string(colorOrange), head, string(colorReset))
	fmt.Printf("Todo Number : %s%d%s\n", string(colorRed), i+1, string(colorReset))
	fmt.Println(out)
}

func PrintAll() {
	fmt.Println(string(colorOrange), head, string(colorReset))
	for i, todo := range db.ReadAll() {
		out, _ := glamour.Render(todo, "dark")
		fmt.Println("--------------")
		fmt.Printf("Todo Number : %s%d%s\n", string(colorRed), i+1, string(colorReset))
		fmt.Println(out)
	}
}

func Edit(number string) {
	if number == "" {
		fmt.Println("please input the a todo number to edit")
		return
	}
	i, err := strconv.Atoi(number)
	i--
	check(err)
	//get from db
	content := db.Read(i)
	// write to file
	err = utils.WriteToFile(filePath, content)
	check(err)
	// edit
	err = utils.EditFile(filePath, editor)
	check(err)
	// read
	content, err = utils.ReadFile(filePath)
	check(err)
	//clear file
	err = utils.WriteToFile(filePath, "")
	check(err)
	// store
	err = db.Replace(i, content)
	check(err)

}

func Add() {
	// make a file
	err := utils.EditFile(filePath, editor)
	check(err)

	// read the file
	content, err := utils.ReadFile(filePath)
	// put it in database
	check(err)
	db.Add(content)
	// empty the file
	err = utils.WriteToFile(filePath, "")

	check(err)
}

func Remove(number string) {
	if number == "" {
		fmt.Println("please input the a todo number to delete")
		return
	}
	i, err := strconv.Atoi(number)
	check(err)

	db.Del(i - 1)
}

func Init(path, editorP string) {
	editor = editorP
	filePath = path + "/todo_temp.md"
}

func Check() bool {
	return db.Size() > 0
}
