package util

import (
	"fmt"
)

type Processor func(string) string

func Test(p Processor, in, res string) {
	fmt.Printf("Testing %s ... ", Type(p))
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

func Trace(p Processor, in string) {
	fmt.Printf("Tracing %s\n", Type(p))
	fmt.Printf("Input:  <%s>\n", in)
	fmt.Printf("Output: <%s>\n", p(in))
}