package main

import (
	"fmt"

	v "github.com/bill-baggins/vector/vector"
)

type Vector2[T comparable] struct {
	x, y T
}

func NewVector2[T comparable](x, y T) Vector2[T] {
	vec := Vector2[T]{}
	vec.x = x
	vec.y = y

	return vec
}

func (vec Vector2[T]) String() string {
	return fmt.Sprint("(", vec.x, ", ", vec.y, ")")
}

func main() {
	vec := v.NewVectorFrom(1, 2, 3, 4, 5, 6)
	fmt.Println(vec.String())

	vecOfVec2 := v.NewVectorFrom(NewVector2(10, 10), NewVector2(10, 20), NewVector2(10, 40))
	fmt.Println(vecOfVec2.String())
}
