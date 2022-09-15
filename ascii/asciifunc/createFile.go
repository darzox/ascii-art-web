package asciifunc

import (
	"fmt"
	"os"
)

func CreateFile(str string) bool {
	f, err := os.Create("files/data.txt")
	if err != nil {
		fmt.Println(err)
		fmt.Println("s")
		return true
	}

	defer f.Close()

	_, err2 := f.WriteString(str + "\n")

	if err2 != nil {
		fmt.Println(err)
		return true
	}
	return false
}
