package testU

import (
    "fmt"
    "../reflectU"
)

type Processor func(string) string

func ExpectOutput(p Processor, in, res string) {
    fmt.Printf("Testing %s ... ", reflectU.TypeName(p))
    out := p(in)
    if out == res {
        fmt.Println("Passed")
    } else {
        fmt.Println("Failed")
        fmt.Printf("Input:  <%s>\n", in)
        fmt.Printf("Output: <%s>\n", out)
        fmt.Printf("Result: <%s>\n", res)
    }
}

func GetOutput(p Processor, in string) {
    fmt.Printf("Running %s\n", reflectU.TypeName(p))
    fmt.Printf("Input:  <%s>\n", in)
    fmt.Printf("Output: <%s>\n", p(in))
}
