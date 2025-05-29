package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("lets play with files")

	filetext := "this goes in a text file"
	filePATH := "./file1.txt"
	file, err := os.Create(filePATH)

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, filetext)

	if err != nil {
		panic(err)
	}

	fmt.Println("lenght: ", length)
	defer file.Close()

	readFile(filePATH)
}


func readFile(filePath string){
	dataInBytes, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}
	fmt.Println("text in file is: ")
	fmt.Println(dataInBytes)

	fmt.Println(string(dataInBytes))
}
