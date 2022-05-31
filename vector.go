// The Vector Library. Contains a single, generic struct: Vector[T].
package main

// Vector is a struct that acts as a wrapper for the slice.
// It contians a lot of helper methods for manipulating the slice.
// It also reduces the boilerplate seen when manipulating slices in the ways
// seen on the SliceTricks GitHub page.
type Vector[T any] struct {
	__array []T
	Length  int
	Cap     int
}

// Creates a new, empty Vector.
// 	Params: None
// 	Returns: Vector[T]
func NewVector[T any]() Vector[T] {
	this := Vector[T]{}
	this.__array = make([]T, 0)
	this.Length = 0
	this.Cap = 0
	return this
}

// Creates a new Vector filled with elements provided in the args parameter.
// 	Params: args ...T
// 	Returns: Vector[T]
func NewVectorFrom[T any](args ...T) Vector[T] {
	this := Vector[T]{}
	this.__array = append(this.__array, args...)
	this.Length = len(this.__array)
	this.Cap = cap(this.__array)
	return this
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
	this.Length += 1
}

// Pops an element from the back of the Vector.
//  Params: None
//  Returns: element T
func (this *Vector[T]) PopBack() T {
	element := this.__array[len(this.__array)-1]
	this.__array = this.__array[:len(this.__array)-1]
	this.Length -= 1
	return element
}

func (this *Vector[T]) PopAt(index int) T {
	value := this.Get(index)
	copy(this.__array[index:], this.__array[index+1:])

	// this.__array[len(this.__array)-1] := nil
	// above can only work with garbage collected types.

	this.__array = this.__array[:len(this.__array)-1]
	this.Length -= 1
	return value
}

// Takes a slice of the Vector.
//  Params: start int
//          end int
//  Returns: Vector[T]
func (this *Vector[T]) Slice(start int, end int) *Vector[T] {
	slice := Vector[T]{}
	slice.__array = this.__array[start:end]
	slice.Length = len(slice.__array)
	slice.Cap = cap(slice.__array)
	return &slice
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
