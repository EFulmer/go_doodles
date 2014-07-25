package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    my_map := make(map[string]int)
    
    for _, word := range strings.Fields(s) {
        /* _, ok := my_map[word]
        
        if !ok {
            my_map[word] = 1
        } else {
            my_map[word] = my_map[word] + 1
        } */
        
        my_map[word] = my_map[word] + 1
    }
    
    return my_map
}

func main() {
    wc.Test(WordCount)
}