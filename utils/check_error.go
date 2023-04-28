package utils

import "fmt"

func CheckError(e error) bool {
	if e != nil {
		fmt.Println("error: ", e)
		return true
	}
	return false
}
