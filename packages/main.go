package main

import (
	"fmt"
	"math"

	"github.com/evgshk/goforit/packages/strutil"
)

func main() {
	fmt.Println(math.Floor(3.6))
	fmt.Println(math.Ceil(3.6))
	fmt.Println(math.Sqrt(25))

	fmt.Println(strutil.Reverse("!haoW"))
}
