// The Vector Library. Contains a single, generic struct: Vector[T].
package vector

import "fmt"

// Vector is a struct that acts as a wrapper for the slice.
// It contians a lot of helper methods for manipulating the slice.
// It also reduces the boilerplate seen when manipulating slices in the ways
// seen on the SliceTricks GitHub page.
type Vector[T any] struct {
	__array []T
}

// Creates a new, empty Vector.
// 	Params: None
// 	Returns: Vector[T]
func NewVector[T any]() Vector[T] {
	this := Vector[T]{}
	this.__array = make([]T, 0)
	return this
}

// Creates a new Vector filled with elements provided in the args parameter.
// 	Params: args ...T
// 	Returns: Vector[T]
func NewVectorFrom[T any](args ...T) Vector[T] {
	this := Vector[T]{}
	this.__array = append(this.__array, args...)
	return this
}

func (this *Vector[T]) Length() int {
	return len(this.__array)
}

func (this *Vector[T]) Cap() int {
	return cap(this.__array)
}

// Returns the element located at the specified index.
//  Params: index int
//  Returns: T element
func (this *Vector[T]) Get(index int) T {
	maybe := this.__array[index]
	return maybe
}

// Change the element at the specified index.
//  Params: index int
//          value T
//  Returns: None
func (this *Vector[T]) Set(index int, value T) {
	this.__array[index] = value
}

// Pushes an element to the back of the Vector.
//  Params: element T
//  Returns: None
func (this *Vector[T]) Push(element T) {
	this.__array = append(this.__array, element)
}

// Pops an element from the back of the Vector.
//  Params: None
//  Returns: element type T
func (this *Vector[T]) PopBack() T {
	element := this.__array[len(this.__array)-1]
	this.__array = this.__array[:len(this.__array)-1]
	return element
}

// Pops an element at a specified position within the Vector.
//	Params: None
// 	Return: T
func (this *Vector[T]) PopAt(index int) T {
	value := this.Get(index)
	copy(this.__array[index:], this.__array[index+1:])

	// this.__array[len(this.__array)-1] := nil
	// above can only work with garbage collected types.
	this.__array = this.__array[:len(this.__array)-1]
	return value
}

// Takes a slice of the Vector.
//  Params: start int
//          end int
//  Return: Vector[T]
func (this *Vector[T]) Slice(start int, end int) *Vector[T] {
	slice := Vector[T]{}
	slice.__array = this.__array[start:end]
	return &slice
}

// Returns the underlying array to be used with the range keyword.
// 	Params: None
//
// 	Return: []T
func (this *Vector[T]) Iter() []T {
	return this.__array
}

// Converts the Vector into a string. Assumes each T type inside of
// the Vector implements the Stringer interface.
// 	Params: None
// 	Return: string
func (this *Vector[Stringer]) String() string {
	strVec := "["
	for i, v := range this.Iter() {
		strVec += fmt.Sprint(v)
		if i != this.Length()-1 {
			strVec += ", "
		}
	}
	strVec += "]"
	return strVec
}

/*
 **           *******       *******     *******    **   ****     **     ********
/**          **/ ////**     **/////**   /**////**  /**  /**/**   /**    **//////**
/**         **     //**   **     //**  /**   /**  /**  /**/ /**  /**   **      //
/**        /**      /**  /**      /**  /*******   /**  /** //** /**  /**
/**        /**      /**  /**      /**  /**/ ///    /**  /**  //**/**  /**    *****
/**        //**     **   //**     **   /**        /**  /**   //****  //**  ////**
/********   //*******     //*******    /**        /**  /**    //***   //********
////////     ///////       ///////     //         //   //      ///     ////////

 ********   **     **   ****     **     ******    **********   **     *******     ****     **    ********
/**/ ////   /**    /**  /**/**   /**    **////**  /////**///   /**    **/////**   /**/**   /**   **//////
/**        /**    /**  /**/ /**  /**   **    //       /**      /**   **     //**  /**/ /**  /**  /**
/*******   /**    /**  /** //** /**  /**             /**      /**  /**      /**  /** //** /**  /*********
/**/ ///    /**    /**  /**  //**/**  /**             /**      /**  /**      /**  /**  //**/**  ////////**
/**        /**    /**  /**   //****  //**    **      /**      /**  //**     **   /**   //****         /**
/**        //*******   /**    //***   //******       /**      /**   //*******    /**    //***   ********
//          ///////    //      ///     //////        //       //     ///////     //      ///   ////////
*/

// Loops through the Vector by value only.
//  Params: actions func(T)
//  Returns: None
func (this *Vector[T]) ForEach(action func(T)) {
	for _, v := range this.__array {
		action(v)
	}
}

// Loops through the Vector in a classic C-style way.
//  Params: start int
//          end int
//          stride int
//          action func(int, T)
//  Returns: None
func (this *Vector[T]) For(start, end, stride int, action func(int, T)) {
	for i := start; i < end; i += stride {
		v := this.__array[i]
		action(i, v)
	}
}

// Uses a function closure that manipulates each Vector element in-place.
//  Params: action func(T) T
//  Returns: None
func (this *Vector[T]) Map(action func(T) T) {
	for i, v := range this.__array {
		this.Set(i, action(v))
	}
}
