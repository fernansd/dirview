package main

import (
    "fmt"
_    "io/fs"
    "log"
    "os"
    "path/filepath"
)

func readDir(path string) ([]string, error) {
    entries, err := os.ReadDir(path)
    if err != nil {
        return nil, err
    }

    var filenames []string
    for _, entry := range entries {
        filenames = append(filenames, entry.Name())
    }

    return filenames, nil
}

func main() {
    targetDir := "."
    absolutePath, _ := filepath.Abs(targetDir)
    fmt.Printf("Listing contents for current directory: %s\n", absolutePath)
    files, err := readDir(targetDir)
    if err != nil  {
        log.Fatal(err)
    }

    for _, file := range files {
        fmt.Println(file)
    }
        

}



