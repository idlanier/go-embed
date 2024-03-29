package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed version.txt
var version string

//go:embed logo.png
var logo []byte

//go:embed files/*
var path embed.FS

func main() {

	fmt.Println(version)

	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)

	if err != nil {
		panic(err)
	}

	dirEntries, _ := path.ReadDir("files")

	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			content, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println("Content:", string(content))
		}
	}
}
