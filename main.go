package main

import (
	"fmt"

	"github.com/sbinet/go-python"
)

func init() {
	if err := python.Initialize(); err != nil {
		panic(err)
	}

	m := python.PyImport_ImportModule("sys")
	if m == nil {
		panic("import error")
	}

	path := m.GetAttrString("path")
	if path == nil {
		panic("get path error")
	}

	currentDir := python.PyString_FromString("")
	python.PyList_Insert(path, 0, currentDir)
}

var (
	pyStr = python.PyString_FromString
	goStr = python.PyString_AsString
	pyInt=python.PyInt_FromLong
	goInt=python.PyInt_AsLong
)

func list() {
	m := python.PyImport_ImportModule("list")
	if m == nil {
		println("module not found")
		return
	}

	persons := m.GetAttrString("list")
	if persons == nil {
		println("persons not found")
		return
	}

	out := persons.CallFunction(pyInt(3))
	if out == nil {
		println("function not found")
		return
	}

	size := python.PyList_Size(out)

	for i := 0; i < size; i++ {
		item := python.PyList_GET_ITEM(out, i)

		name:=item.GetAttrString("name")
		age:=item.GetAttrString("age")

		fmt.Printf("golang: name's: %s, age: %d \n", goStr(name), goInt(age))
	}
}

func main() {
	list()
}
