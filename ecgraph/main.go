package main

import (
        "flag"
        "fmt"
        "math"
)

func main() {

        var str string  // store points lists as a string
        var p, a, b int // variables
        var csv bool    // csv?

        flag.IntVar(&p, "p", 11, "modulus")
        flag.IntVar(&a, "a", -2, "'a' value")
        flag.IntVar(&b, "b", 0, "'b' value")
        flag.BoolVar(&csv, "csv", false, "output points as CSV at the end")
        flag.Parse()

        if !csv {
                str = "(inf)"
        }

        points := 1

        fmt.Printf("\nE: y^2 = x^3 + %dx + %d (mod %d)\n\n", a, b, p)
        for y := p - 1; y >= 0; y-- {
                for x := 0; x < p; x++ {
                        x_cubed := int(math.Pow(float64(x), 3))
                        if ((p*p-y*y)+x_cubed+a*x+b)%p == 0 {
                                points++
                                if csv {
                                        str = fmt.Sprintf("%d, %d\n%s", x, y, str)
                                } else {
                                        str = fmt.Sprintf("(%d, %d) %s", x, y, str)
                                }
                                print("@ ")
                        } else {
                                print(". ")
                        }
                }
                fmt.Printf("%d\n", y)
        }

        // add an extra newline at the beginning if we're outputting csv
        if csv {
                str = fmt.Sprintf("\n%s", str)
        }

        fmt.Printf("\nE= %s\n|E|= %d\n\n", str, points)
}
