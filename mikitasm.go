package main

import (
	"fmt"
	"os"
)

func main() {
	data := make([]byte, 100)
	file, err := os.OpenFile("testdata.txt", os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("[DEBUG] Error in opening file")
		fmt.Println("[DEBUG] Trying to create new file")
		CreateNewFile("test.txt")
	} else {
		fmt.Println("[DEBUG] File opened successfully", &file)
	}
	ReadFromFile(data, file)

	//part with writing data to file dont fucking work, idk for now how to fix it, but I got new data in file with data
	_, err = fmt.Fprintln(file, data)
	if err != nil {
		fmt.Println("[DEBUG] PIZDEZ DOESNT WORK!!!", err)
		file.Close()
		return
	} else {
		fmt.Println("[DEBUG] All is good")
	}
	fmt.Println("[DEBUG] Trying to read data form modified file")
	ReadFromFile(data, file)
	file.Close()
}

//this func does work correctly
func CreateNewFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("[DEBUG] Error with creating new file")
		fmt.Println("[DEBUG] Stopping program")
	} else {
		fmt.Println("[DEBUG] File created successfully", &file)
	}
	return file
}

//this func does work correctly, but need to be modified
func ReadFromFile(data []uint8, file *os.File) {
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
