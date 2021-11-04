package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/Ghibranalj/todo-cli/modes"
)

//$HOME/.todo-cli/ideas
//$HOME/.todo-cli/todos
// can edit with editor of choise
// editor saven in $HOME/

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

	args := os.Args[:1]

	switch strings.ToLower(args[0]) {
	case "todo", "todos":
		modes.Todo(args[1])
	case "idea", "ideas":
		modes.Idea(args[1])
	case "help":
		help()
	default:
		fmt.Printf("action %s not found \n", args[0])
		help()
	}
}

func help() {

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

	modes.Editor = Conf.Editor
}

func askForEditor() string {

	posEditors := []string{
		"nano", "nvim", "vim", "emacs",
	}
	editors := []string{}
	for _, posEditor := range posEditors {
		out, _ := exec.Command("/usr/bin/which", posEditor).Output()
		outS := string(out[:])
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
		fmt.Printf("[%d] %s", i+1, editor)
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
