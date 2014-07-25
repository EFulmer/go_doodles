package main

import "fmt"

func a() int {
        i := 0
        defer fmt.Println(i)
        i++
        return i
}

func main() { 
        v := a() 
        fmt.Println(v)
}
