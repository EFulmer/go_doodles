package main

import (
    "code.google.com/p/go-tour/pic"
    "math"
)

func Pic(dx, dy int) [][]uint8 {
    res := make([][]uint8, dy)
    for i := range res {
        c := make([]uint8, dx)
        for j := range c {
            c[j] = uint8( math.Sin(float64(i*j)) * 10 )
        }
        
        res[i] = c
    }
    
    return res
}

func main() {
    pic.Show(Pic)
}