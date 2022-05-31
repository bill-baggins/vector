package lib

type Vector[T any] struct {
	__array []T
	Length  int
	Cap     int
}

func NewVector[T any]() Vector[T] {
	this := Vector[T]{}
	this.__array = make([]T, 0)
	this.Length = 0
	this.Cap = 0
	return this
}

func NewVectorFrom[T any](args ...T) Vector[T] {
	this := Vector[T]{}
	this.__array = append(this.__array, args...)
	this.Length = len(this.__array)
	this.Cap = cap(this.__array)
	return this
}

func (this *Vector[T]) Get(index int) T {
	maybe := this.__array[index]
	return maybe
}

func (this *Vector[T]) Set(index int, value T) {
	this.__array[index] = value
}

func (this *Vector[T]) Push(element T) {
	this.__array = append(this.__array, element)
	this.Length += 1
}

func (this *Vector[T]) PopBack() T {
	element := this.__array[len(this.__array)-1]
	this.__array = this.__array[:len(this.__array)-1]
	this.Length -= 1
	return element
}

func (this *Vector[T]) ForEach(action func(T)) {
	for _, v := range this.__array {
		action(v)
	}
}

func (this *Vector[T]) For(start int, end int, stride int, action func(int, T)) {
	for i := start; i < end; i += stride {
		v := this.__array[i]
		action(i, v)
	}
}

func (this *Vector[T]) Map(action func(T) T) {
	for i, v := range this.__array {
		this.Set(i, action(v))
	}
}
