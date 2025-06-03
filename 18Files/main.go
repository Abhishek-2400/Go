package main

import (
	"fmt"
	"io"
	"os"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readMyFile(filePath string) {
	databyte, err := os.ReadFile(filePath) //earlier ioutil.ReadFile was used, but os.ReadFile is more idiomatic in Go 1.16+
	checkError(err)
	fmt.Printf("Data read from file: %v", string(databyte))
}

func main() {
	fmt.Println("Files in Go")

	//file creation
	file, err := os.Create("./example.txt")
	checkError(err)
    

	//file writing
	length, err := io.WriteString(file, "Hello, this is a test file.\n")
	checkError(err)
	fmt.Printf("Length of written data: %d bytes\n", length)

	defer file.Close() // may be after this u want to write more code , // so defer will close the file at the end of main function execution

	//file reading
	readMyFile("./example.txt")
}

//os module is used for file creation 
//io module is used for writing data to file
//os.ReadFile is used to read the file

