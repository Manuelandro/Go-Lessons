package fileutils

import "fmt"

func IsError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}