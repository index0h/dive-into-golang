package main

import "fmt"

func main() {
    trueVar := true
    falseVar := false
    
    fmt.Print("First: ")
    
    if trueVar {
        fmt.Println("if trueVar")
    }
    
    fmt.Print("First: ")
    
    if falseVar {
        fmt.Println("Second: if falseVar")
    } else {
        fmt.Println("Second: if falseVar -> else")
    }
}
