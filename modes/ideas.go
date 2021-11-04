package modes

func Idea(args []string) {
	switch args[0] {
	case "print", "show":
		print("ideas")
	case "edit", "write":
		edit("ideas")
	}
}
