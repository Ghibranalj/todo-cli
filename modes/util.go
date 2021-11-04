package modes

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

var Editor string
var Path string

const (
	todoHead = ` _______   ____    _____     ____  
|__  __|  / __ \  |  __ \   / __ \ 
  | |    | |  | | | |  | | | |  | |
  | |    | |  | | | |  | | | |  | |
  | |    | |__| | | |__| | | |__| |
  |_|     \____/  |_____/   \____/ `

	ideaHead = `_____    _____    ______               _____ 
|_   _| |  __  \ |  ____|     /\      / ____|
  | |   | |  | | | |__       /  \    | (___  
  | |   | |  | | |  __|     / /\ \    \___ \ 
 _| |_  | |__| | | |____   / ____ \   ____) |
|_____| |_____/  |______| /_/    \_\ |_____/ 
`
)

func check(e error) {
	if e != nil {
		fmt.Printf("There is an Error: %s \n", e.Error())
		os.Exit(1)
	}
}

func print(mode string) {

	filename := ""
	head := ""

	switch mode {
	case "todos":
		filename = "todos"
		head = todoHead

	case "ideas":
		filename = "ideas"
		head = ideaHead
	default:
		return
	}

	file, err := ioutil.ReadFile(Path + "/" + filename)
	check(err)
	fmt.Println(head)
	fmt.Println(string(file[:]))
}

func edit(mode string) {
	cmd := exec.Command(Editor, Path+"/"+mode)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	check(err)
}

func Init() {
	if _, err := os.Stat(Path + "/todos"); os.IsNotExist(err) {
		err = os.WriteFile(Path+"/todos", []byte(""), 0644)
		check(err)
	}

	if _, err := os.Stat(Path + "/ideas"); os.IsNotExist(err) {
		err = os.WriteFile(Path+"/ideas", []byte(""), 0644)
		check(err)
	}

}
