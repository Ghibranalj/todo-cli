package modes

func Todo(args []string) {
	switch args[0] {
	case "print", "show":
		print("todos")
	case "edit", "write":
		edit("todos")
	}
}
