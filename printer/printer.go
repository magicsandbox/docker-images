package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {

	fileLocation := os.Getenv("FILE_DIRECTORY")
	var filePath string

	if (fileLocation != "") {
		filePath = fileLocation + "/replaceme.txt"
	} else {
		filePath = "./replaceme.txt"
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		data := []byte("The contents of this file should be replaced")
		ioutil.WriteFile(filePath, data, 0644)
	}
	for {
		dat, err := ioutil.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dat))
		<-time.After(1 * time.Second)
	}
}
