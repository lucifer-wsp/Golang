package main

import (
    "fmt"
)

func pointer_array_mod(){
    s := [3]int{1,2,3}
    fmt.Println(s)
    func (ar *[3]int){
        ar[0] = 4
        fmt.Println(ar)
    }(&s)
    fmt.Println(s)
}

func copy_array_mod(){
    s := [3]int{1,2,3}
    fmt.Println(s)
    func (ar [3]int){
        ar[0] = 4
        fmt.Println(ar)
    }(s)
    fmt.Println(s)
}

func main(){
    pointer_array_mod()
    copy_array_mod()
}