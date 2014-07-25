package main

import (
    "fmt"
    "math"
)

const Epsilon = 0.000000000001

func Sqrt(x float64) float64 {
    var num, denom float64
    
    z := x/2
    for i := 0; i < 10; i++ {
        num = math.Pow(z, 2.0) - x
        denom = 2 * x
        z = z - (num / denom)
    }
    
    return z
}

func SqrtNew(x float64) float64 {
    var num, denom float64
    
    z := x/2
    
    for delta := x; math.Abs(delta - z) > Epsilon; {
        delta = z
        num = math.Pow(z, 2.0) - x
        denom = 2 * x
        z = z - (num / denom)
    }
    
    return z
}

func main() {
    fmt.Println(SqrtNew(2))
}