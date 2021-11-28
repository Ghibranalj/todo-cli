package modes

import (
	"fmt"

	"github.com/charmbracelet/glamour"
)

// TODO use markdown
func Help() {
	helpmd := `
todo-cli by Ghibranalj
# Usage :
## todo-cli [command] (todo)
	commands:
	add    : add a todo to your todo list
	edit   : edit a todo in your todo list
	print  : print a todo in your todo list or print all if (todo) is empty
	remove : remove a todo in your todo list
	help   : show this message
	check  : $? will be 0 if there is todo in todo list
	bashrc : print bash script for .bashrc`

	out, err := glamour.Render(helpmd, "dark")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(out)
}
