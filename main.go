package main

import "fmt"

// https://stackoverflow.com/questions/28081486/how-can-i-go-run-a-project-with-multiple-files-in-the-main-package

func main() {

	//Goals:
	// 1. Read and write data into a DB file
	// 2. Store the data in the file in a structured way ???
	// 3. REPL interface

	err := Repl()
	if err != nil {
		fmt.Println(err)
	}

	/*filename := "test"
	data := "Hello World"

	// Write
	{
		f, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		f.WriteAt([]byte(data), 0)
		f.Close()
	}

	// Read
	{
		f, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		d := make([]byte, 5)
		f.ReadAt(d, 6)
		fmt.Println(string(d))
	}*/

}
