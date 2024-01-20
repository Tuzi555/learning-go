package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		b      byte  = math.MaxUint8
		smallI int32 = math.MaxInt32
		bigI   int64 = math.MaxInt64
	)

	b += 1
	smallI += 1
	bigI += 1

	fmt.Printf("b: %d, smallI: %d, bigI: %d", b, smallI, bigI)
}
