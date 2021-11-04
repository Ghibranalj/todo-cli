package modes

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/charmbracelet/glamour"
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

func Print(mode string) {

	filename := ""
	head := ""

	switch mode {
	case "todos", "todo":
		filename = "todos"
		head = todoHead

	case "ideas", "idea":
		filename = "ideas"
		head = ideaHead
	default:
		fmt.Println("What do you want to print")
		fmt.Println("[1] Todos")
		fmt.Println("[2] Ideas")
		filenames := []string{"todos", "ideas"}
		i := 1
		fmt.Scanf("%d", &i)
		head = todoHead
		if i == 2 {
			head = ideaHead
		}
		filename = filenames[i-1]
	}

	file, err := ioutil.ReadFile(Path + "/" + filename)
	check(err)
	fmt.Println(head)
	// fmt.Println(string(file[:]))
	out, err := glamour.Render(string(file[:]), "dark")
	check(err)
	fmt.Print(out)
}

func Edit(mode string) {
	filename := ""
	switch mode {
	case "todos", "todo":
		filename = "todos"

	case "ideas", "idea":
		filename = "ideas"
	default:
		fmt.Println("What do you want to edit")
		fmt.Println("[1] Todos")
		fmt.Println("[2] Ideas")
		filenames := []string{"todos", "ideas"}
		i := 1
		fmt.Scanf("%d", &i)
		filename = filenames[i-1]
	}
	cmd := exec.Command(Editor, Path+"/"+filename)
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
