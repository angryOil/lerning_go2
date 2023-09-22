package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	toFile := person{
		Name: "joy",
		Age:  11,
	}

	tmpFile, err := os.CreateTemp(os.TempDir(), "sample-")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())
	err = json.NewEncoder(tmpFile).Encode(toFile)
	if err != nil {
		panic(err)
	}
	err = tmpFile.Close()
	if err != nil {
		panic(err)
	}

	tmpFile2, err := os.Open(tmpFile.Name())
	if err != nil {
		panic(err)
	}

	var fromeFile person
	err = json.NewDecoder(tmpFile2).Decode(&fromeFile)
	if err != nil {
		panic(err)
	}
	err = tmpFile2.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", fromeFile)
}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
