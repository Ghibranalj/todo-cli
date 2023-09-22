package utils

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func ReadFile(path string) (string, error) {

	bytes, err := ioutil.ReadFile(path)
	return string(bytes), err
}

func WriteToFile(path, content string) error {
	err := os.WriteFile(path, []byte(content), 0644)

	return err
}

func EditFile(path, editor string) error {


	// run bash -c "editor path"
	cmd := exec.Command("bash", "-c", editor+" "+path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	return err
}

// CreateFile creates file if file DNE
func CreateFile(path, content string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.WriteFile(path, []byte(content), 0644)
	}
	return err
}
