package main

import (
	"os"
	"fmt"
	"io"
)

func main() {
	// use Args comming from terminal to specify the filename.
	readFile(os.Args[1])
}

/*
 * this function gets the filename by argument
 * type of string, and gets all content from the file
 * using Copy from IO package, I left commented out
 * another approach I wrote that works fine as well.
 */
func readFile(fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)

	// data := make([]byte, 100)
	// count, err := file.Read(data)
	// output := string(data[:count])
	// fmt.Println(output)

}
