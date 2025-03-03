package main

import (
	"fmt"
	"os"
	"io/ioutil"

    "path/filepath"
)

func main() {
	// files, err := ioutil.ReadDir("../.")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// }

	currentPath, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

	fmt.Println(currentPath)
    // Пытаемся подняться до корня
    var lastValidPath string
    for {
        _, err = ioutil.ReadDir(currentPath)
        if err != nil {
            break
        }
        
        lastValidPath = currentPath
        currentPath = 	filepath.Dir(currentPath) 
        fmt.Println(currentPath)
        if currentPath == filepath.Dir(currentPath) {
            break
        }
    }

    files, err := ioutil.ReadDir(lastValidPath)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    
    fmt.Println("Root directory:", lastValidPath)
    for _, file := range files {
        fmt.Println(file.Name())
    }
}
