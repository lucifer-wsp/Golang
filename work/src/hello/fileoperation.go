package main

import (
    "fmt"
    "os"
)

func main(){
    filepath := "/opt/lucifer/Golang/bin/hello"
    if _, err := os.Stat(filepath); os.IsNotExist(err) {
        fmt.Println("file is already exist !")
    }else{
        fmt.Println("file not found !")
    }
}