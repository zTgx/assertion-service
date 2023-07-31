package main

import (
	"fmt"
	"strings"
	"os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func FilterNames(message string) string {
	filterString := strings.ReplaceAll(message, "should be one of the following; ", "")
	return filterString
}

func ParseMockNames(message string) []string {
	res := strings.Split(message, ",")
	fmt.Println("[split] Achainable Names Count: ", len(res))

	return res
}

func WriteToFile(path string, content string) {
	f, err := os.Create(path)
	check(err)

	defer f.Close()

	n, err := f.WriteString(content)
    check(err)
    fmt.Printf("wrote %d bytes\n", n)

	f.Sync()
}