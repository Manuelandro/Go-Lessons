package fileutils

import (
	"fmt"
	"os"
)

func WriteFile(path string, content string) {
	// open file using READ & WRITE permission
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if IsError(err) {
		return
	}
	defer file.Close()

	// write some text line-by-line to file
	_, err = file.WriteString(content)
	if IsError(err) {
		return
	}

	// save changes
	err = file.Sync()
	if IsError(err) {
		return
	}

	fmt.Println("==> file scritto: ", path)
}
