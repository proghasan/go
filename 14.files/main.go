package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

const FileNameWithPath = "./myContent.txt"

func main() {
	fmt.Println("Welcome to files")

	content := "We learning golang. current lesson files."
	file, err := os.Create(FileNameWithPath)

	// close the
	defer file.Close()

	// error catch
	checkNilError(err)

	// write content into file
	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("Content length is: ", length)

	// read file
	readFile(FileNameWithPath)

}

func readFile(fileName string) {
	dataByte, err := ioutil.ReadFile(fileName)
	checkNilError(err)
	fmt.Println("Content reading:  ", string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
