package main

import (
	"fmt"
	"os"
)

func main() {
	data := make([]byte, 100)
	//var current_filepath string = "C:/development/GO/main/test.txt"
	file, err := os.Open("C:/development/GO/main/test.txt")
	if err != nil {
		fmt.Println("[DEBUG] Error in opening file")
		fmt.Println("[DEBUG] Trying to create new file")
		create_new_file("test.txt")
	} else {
		fmt.Println("[DEBUG] File opened successfully", &file)
	}
	read_from_file(data, file)
	file.Close()
}

func create_new_file(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("[DEBUG] Error with creating new file")
		fmt.Println("[DEBUG] Stopping program")
	} else {
		fmt.Println("[DEBUG] File created successfully", &file)
	}
	return file
}

func read_from_file(data []uint8, file *os.File) {
	count, err := file.Read(data)
	if err != nil {
		//if file is empty it return error, we need to fix it (need to return warning, and firstly recognize it somehow)
		fmt.Println("[DEBUG] Error in reading data from file")
		fmt.Println("[DEBUG] Stopping program")
	} else {
		fmt.Println("[DEBUG] Reading success", &file)
		fmt.Printf("read %d bytes: %q\n", count, data[:count])
	}
}
