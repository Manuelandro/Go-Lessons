package fileutils

import (
	"fmt"
	"os"
)

func CreateFile(path string, content string) {

	var _, err = os.Stat(path)

	if os.IsNotExist(err) {
		file, err := os.Create(path)
		if IsError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("==> file creato: ", path)
	WriteFile(path, content)

}
