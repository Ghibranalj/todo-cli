package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

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
if [ $? ]
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

	if _, err := os.Stat(Path + "/conf.json"); os.IsNotExist(err) {
		conf := conf{
			Editor: askForEditor(),
		}
		confText, _ := json.Marshal(conf)

		err = os.WriteFile(Path+"/conf.json", []byte(confText), 0644)
		check(err)

	}
	byteConf, err := ioutil.ReadFile(Path + "/conf.json")
	check(err)

	err = json.Unmarshal(byteConf, &Conf)
	check(err)

	err = db.Init(Path)
	check(err)
	modes.Init(Path, Conf.Editor)
}

func askForEditor() string {

	posEditors := []string{
		"nano", "nvim", "vim", "emacs",
	}
	editors := []string{}
	for _, posEditor := range posEditors {
		out, _ := exec.Command("/usr/bin/which", posEditor).Output()
		outS := strings.TrimSuffix(string(out[:]), "\n")
		if outS == "" {
			continue
		}
		editors = append(editors, outS)
	}

	if len(editors) == 0 {
		fmt.Printf("There is an Error: %s \n", "no editor can be found")
		os.Exit(1)
	}
	for i, editor := range editors {
		fmt.Printf("[%d] %s \n", i+1, editor)
	}
	i := 1
	fmt.Print("Select your editor [default:1]")
	_, err := fmt.Scanf("%d", &i)
	if err != nil || i > len(editors) {
		fmt.Println("Default value (1) selected")
		i = 1
	}
	return editors[i-1]
}
