package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Ghibranalj/todo-cli/db"
	"github.com/Ghibranalj/todo-cli/modes"
)

//$HOME/.todo-cli/ideas
//$HOME/.todo-cli/todos
// can edit with editor of choise
// editor saved in $HOME/

type conf struct {
	Editor string `json:"editor"`
}

var Path string
var Conf conf

func check(e error) {
	if e != nil {
		fmt.Printf("There is an Error: %s \n", e.Error())
		os.Exit(1)
	}
}
func main() {

	args := os.Args[1:]

	if len(args) == 0 {
		args = append(args, "help")
	}

	command, option := args[0], ""

	if len(args) >= 2 {
		option = args[1]
	}

	switch command {

	case "reset":
		modes.Reset()
	case "add":
		modes.Add()
	case "edit":
		modes.Edit(option)
	case "print":
		if option == "" {
			modes.PrintAll()
		} else {
			modes.Print(option)
		}
	case "check":
		if modes.Check() {
			os.Exit(0)
		}
		os.Exit(1)
	case "bashrc":
		out := `todo-cli check 2> /dev/null
if [ $? -eq 0 ] && [ "$PWD" == "$HOME" ]
then
	todo-cli print
fi`
		fmt.Println(out)
	case "remove", "delete":
		modes.Remove(option)
	case "", "help":
		modes.Help()
	default:
		fmt.Printf("Command %s not valid\n", command)
		modes.Help()
	}

	db.Save()
}

func init() {
	home, err := os.UserHomeDir()
	check(err)

	Path = home + "/.todo-cli"
	if _, err := os.Stat(Path); os.IsNotExist(err) {
		err := os.Mkdir(Path, 0700)
		check(err)
	}

	editor := os.Getenv("EDITOR")

	for editor == "" {
		// ask for editor
		fmt.Println("Please enter your editor of choise")
		fmt.Scanln(&editor)
	}

	if _, err := os.Stat(Path + "/conf.json"); os.IsNotExist(err) {
		conf := conf{
			Editor: editor,
		}
		confText, _ := json.Marshal(conf)

		err = os.WriteFile(Path+"/conf.json", []byte(confText), 0644)
		check(err)

	}
	byteConf, err := os.ReadFile(Path + "/conf.json")
	check(err)

	err = json.Unmarshal(byteConf, &Conf)
	check(err)

	err = db.Init(Path)
	check(err)
	modes.Init(Path, Conf.Editor)
}
