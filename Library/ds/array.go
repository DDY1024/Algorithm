package main

import "fmt"

type Array struct {
	values []interface{}
}

func New(size int) *Array {
	return &Array{values: make([]interface{}, size)}
}

func NewFromArray(other *Array) *Array {
	a := &Array{values: make([]interface{}, other.Size())}
	for i := range other.values {
		a.values[i] = other.values[i]
	}
	return a
}

func (a *Array) Fill(val interface{}) {
	for i := range a.values {
		a.values[i] = val
	}
}

func (a *Array) Set(pos int, val interface{}) {
	if pos < 0 || pos >= len(a.values) {
		return
	}
	a.values[pos] = val
}

func (a *Array) At(pos int) interface{} {
	if pos < 0 || pos >= len(a.values) {
		return nil
	}
	return a.values[pos]
}

func (a *Array) Front() interface{} {
	return a.At(0)
}

func (a *Array) Back() interface{} {
	return a.At(len(a.values) - 1)
}

func (a *Array) Size() int {
	return len(a.values)
}

func (a *Array) Empty() bool {
	return a.Size() == 0
}

func (a *Array) SwapArray(other *Array) {
	if a.Size() != other.Size() {
		return
	}
	a.values, other.values = other.values, a.values
}

func (a *Array) Data() []interface{} {
	return a.values
}

func (a *Array) String() string {
	return fmt.Sprintf("%v", a.values)
}
