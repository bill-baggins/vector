package main

import (
	"testing"
)

type T = interface{}

func TestNewVector(t *testing.T) {
	slice := make([]T, 0)
	vec := NewVector[T]()
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
	if myVec.Length != expected {
		t.Log("(Vector).Push Method: FAIL")
		t.Log("Expected: myVec.Length == 5\nGot: myVec.Length == 4\n")

		t.Fail()
	}
}

func TestPopBack(t *testing.T) {
	expected := 3
	myVec := NewVectorFrom(1, 2, 3, 4)
	myVec.PopBack()
	if myVec.Length != expected {
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
