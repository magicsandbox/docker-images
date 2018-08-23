package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {

	for {
		dat, err := ioutil.ReadFile("./replaceme.txt")
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dat))
		<-time.After(1 * time.Second)
	}
}
