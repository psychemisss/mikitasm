package main

import (
	"fmt"
	"os"
)

func main() {
	data := "fd"
	file, err := os.OpenFile("testdata.txt", os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("[DEBUG] Error in opening file :: ", err)
		fmt.Println("[DEBUG] Trying to create new file")
		CreateNewFile("test.txt")
	} else {
		fmt.Println("[DEBUG] File opened successfully", &file)
	}
	ReadFromFile(data, file)

	//part with writing data to file dont fucking work, idk for now how to fix it, but I got new data in file with data
	WriteToFile(data, file)
	fmt.Println("[DEBUG] Trying to read data form modified file")
	ReadFromFile(data, file)

	file.Close()
}

func WriteToFile(data string, file *os.File) {
	file.Write([]byte(data))
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
func ReadFromFile(data string, file *os.File) {
	//getting file stats, and whis stuff fking work
	filestatus, err := os.Stat("testdata.txt")
	if err != nil {
		fmt.Println("[DEBUG] Can't get file status :: ", err)
		fmt.Println("[DEBUG] Stopping program")
		file.Close()
	} else {
		fmt.Println("[DEBUG] File content size is", filestatus.Size())
	}

	if filestatus.Size() == 0 {
		fmt.Println("[DEBUG] There is no content in file, can't read anything")
		file.Close()
		os.Exit(1)
	}

	//if file content size is 0, reading doesn't work
	data1 := make([]byte, 100)
	count, err := file.ReadAt(data1, 0)
	data = string(data1)
	fmt.Println("[DEBUG] Reading success", &file)
	fmt.Printf("read %d bytes: %q\n", count, data[:count])

}