package main

import (
        "flag"
        "fmt"
        "math"
)

func main() {

        var p, a, b int

        flag.IntVar(&p, "p", 11, "modulus")
        flag.IntVar(&a, "a", -2, "'a' value")
        flag.IntVar(&b, "b", 0, "'b' value")
        flag.Parse()

        str := "(inf)"
        points := 1

        fmt.Printf("E: y^2 = x^3 + %dx + %d (mod %d)\n", a, b, p)
        for y := p - 1; y >= 0; y-- {
                x := -1
                for x = 0; x < p; x++ {
                        x_cubed := int(math.Pow(float64(x), 3))
                        if ((p*p-y*y)+x_cubed+a*x+b)%p == 0 {
                                points++
                                str = fmt.Sprintf("(%d, %d) %s", x, y, str)
                                print("@ ")
                        } else {
                                print(". ")
                        }
                }
                fmt.Printf("%d\n", y)
        }
        fmt.Printf("E= %s\n|E|= %d\n", str, points)
}
