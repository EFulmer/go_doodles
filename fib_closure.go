package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    a, b := 1, 1
    my_f := func() int {
        a, b = b, a+b
        return a
    }
    
    return my_f
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}