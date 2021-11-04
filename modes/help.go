package modes

import "fmt"

func Help() {
	fmt.Println(`
TODO-cli by Ghibranalj.
Usage : todo-cli [mode] [target]

Target: 
	todos, todo : your todos.
	idea, ideas : your ideas.
Modes:
	print, show : print [target] to STDOUT.
	edit, write : edit [target] using selected editor.`)
}
