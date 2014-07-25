package main

import (
    "fmt"
    "math/cmplx"
)

const epsilon = 0.000000001

func Cbrt(x complex128) complex128 {
    z := x / 2
    
    for delta := x; cmplx.Abs(delta - z) > epsilon; {
        delta = z
        num := cmplx.Pow(z, 3.0) - x
        denom := 3.0 * cmplx.Pow(z, 2.0)
        z = z - (num / denom)
    }
    
    return z
}

func main() {
    fmt.Println(Cbrt(27))
}