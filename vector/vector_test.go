package vector

import (
	"testing"
)

func TestNewVector(t *testing.T) {
	slice := make([]any, 0)
	vec := NewVector[any]()
	if len(vec.__array) != len(slice) {
		t.Log("Array not initialized correctly\n")
		t.Log("Expected: Array length == 0")
		t.Log("Got: undefined array length")
		t.Fail()
	}
}

func TestPush(t *testing.T) {
	expected := 5
	myVec := NewVectorFrom(1, 2, 3, 4)
	myVec.Push(5)
	if myVec.Length() != expected {
		t.Log("(Vector).Push Method: FAIL")
		t.Log("Expected: myVec.Length == 5\nGot: myVec.Length == 4\n")

		t.Fail()
	}
}

func TestPopBack(t *testing.T) {
	expected := 3
	myVec := NewVectorFrom(1, 2, 3, 4)
	myVec.PopBack()
	if myVec.Length() != expected {
		t.Log("Expected: myVec.Length == 3\nGot: myVec.Length == 4\n")
		t.Log("(Vector).PopBack Method: FAIL")
		t.Fail()
	}
}

func TestGet(t *testing.T) {
	expected := "Frank"
	myVec := NewVectorFrom("Bob", "Jill", "Frank", "Curtis")
	got := myVec.Get(2)
	if got != expected {
		t.Log("(Vector).Get Method: FAIL")
		t.Log("Expected: Frank")
		t.Logf("Got: %s", got)
		t.Fail()
	}
}

func TestSet(t *testing.T) {
	expected := "Hailey"
	myVec := NewVectorFrom("Bob", "Jill", "Frank", "Curtis")
	myVec.Set(2, "Hailey")
	got := myVec.Get(2)
	if got != expected {
		t.Log("(Vector).Set Method: FAIL")
		t.Log("Expected: Hailey")
		t.Logf("Got: %s", got)
		t.Fail()
	}
}

type person struct {
	name string
	age  int
}

func TestPopAt(t *testing.T) {
	// Testing with 'primitive' types.
	expected := []int{1, 2, 4, 5, 6}
	myVec := NewVectorFrom(1, 2, 3, 4, 5, 6)
	myVec.PopAt(2) // should remove the third element.
	areEqual := true
	myVec.For(0, myVec.Length(), 1, func(i int, v int) {
		v2 := expected[i]
		if v2 != v {
			areEqual = false
			return
		}
	})
	if !areEqual {
		t.Log("(Vector).PopAt Method: FAIL")
		t.Log("Expected: [1, 2, 4, 5, 6]")
		t.Log("Got: ", myVec.__array)
		t.Fail()
	}

	// Garbage collected type test.
	peopleExpected := []*person{
		{"Hailey", 19},
		{"Bob", 23},
		{"Yumi", 20},
		{"Mimi", 45},
	}

	peopleVec := NewVectorFrom(
		&person{"Hailey", 19},
		&person{"Bob", 23},
		&person{"Yumi", 20},
		&person{"Curtis", 34},
		&person{"Mimi", 45},
	)

	_ = peopleVec.PopAt(peopleVec.Length() - 2)
	_ = peopleVec.PopAt(peopleVec.Length() - 2)
	_ = peopleVec.PopAt(peopleVec.Length() - 2)

	areEqual = true
	peopleVec.For(0, peopleVec.Length(), 1, func(i int, v *person) {
		v2 := peopleExpected[i]
		if v.name != v2.name || v.age != v2.age {
			areEqual = false
			return
		}
	})

	if !areEqual {
		t.Log("(Vector).PopAt Method: FAIL")
		t.Log("Expected: ", peopleExpected)
		t.Log("Got: ", peopleVec.__array)
		t.Fail()
	}
}
