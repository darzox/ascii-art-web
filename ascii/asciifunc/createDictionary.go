// Creating map of ascii symbols, where key is decimal in ASCII and value is symbol

package asciifunc

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
)

var Dict map[byte]string = map[byte]string{}

func CreateDict(fontTemplate string) bool {
	pathOfTemplate := "ascii/templates/" + fontTemplate + ".txt"
	hashesOfTemplates := map[string]string{
		"banner":     "6581cab5cd8daadf23a72d4867fcb39a",
		"standard":   "ac85e83127e49ec42487f272d9b9db8b",
		"shadow":     "a49d5fcb0d5c59b2e77674aa3ab8bbb1",
		"thinkertoy": "bf1d925662e40f5278b26a0531bfdb63",
	}
	// checks for changes in template files
	if hashesOfTemplates[fontTemplate] != md5sum(pathOfTemplate) {
		return false
	}
	file, err := os.Open(pathOfTemplate)
	if err != nil {
		return false
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	afterFirstline := false
	letter := ""
	i := 32
	for scanner.Scan() {
		if len(scanner.Text()) == 0 && !afterFirstline {
			afterFirstline = true
			continue
		}
		if afterFirstline {
			letter += scanner.Text() + "\n"
		}
		if len(scanner.Text()) == 0 && afterFirstline {
			for i <= 126 {
				Dict[byte(i)] = letter
				break
			}
			letter = ""
			i++
			continue
		}
	}
	nLineInLetter := 9
	// adds last letter to dictionary
	letter += "\n"
	Dict[126] = letter
	// checks for banner format
	for _, l := range Dict {
		count := 0
		for _, s := range l {
			if s == '\n' {
				count++
			}
		}
		if count != nLineInLetter {
			return false
		}
	}
	return true
}

func md5sum(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return ""
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return ""
	}
	return hex.EncodeToString(hash.Sum(nil))
}
