package oden

import (
	"fmt"
	"github.com/willf/bitset"
	"reflect"
)

type Object struct {
	data map[uint]IData
	bits bitset.BitSet
}

func NewObject() *Object {
	return &Object{
		data: make(map[uint]IData),
	}
}

func (this *Object) AddData(data IData) {
	// wop wop
	this.data[dataManager.Get(data)] = data
	this.bits.Set(dataManager.Get(data))
}

func (this *Object) Bits() *bitset.BitSet {
	return &this.bits
}

func (this *Object) DataByName(name string) IData {
	for _, data := range this.data {
		if reflect.TypeOf(data).String() == name {
			return data
		}
	}
	return nil
}

func (this *Object) DataByType(data IData) IData {
	datatype := reflect.TypeOf(data)
	for _, data := range this.data {
		if datatype == reflect.TypeOf(data) {
			return data
		}
	}

	return nil
}

func (this *Object) DataByIndex(index uint) IData {
	return this.data[index]
}

func (this *Object) DebugData() {
	for k, data := range this.data {
		fmt.Println("[", k, "]", reflect.TypeOf(data).String())
	}
}
