package main

import (
	"fmt"

	l "github.com/bill-baggins/vector/lib"
)

func main() {
	myVec := l.NewVectorFrom(1, 2, 3, 4, 5, 6)
	myVec.Push(7)
	myVec.Push(8)
	myVec.PopBack()
	myVec.PopBack()

	myVec.Map(func(v int) int {
		return v / 2
	})

	myVec.ForEach(func(v int) {
		fmt.Printf("%d\n", v)
	})

}
