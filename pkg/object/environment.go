package object

import (
	"bytes"
	"io"
)

type Environment struct {
	store  map[string]Object
	outer  *Environment
	output io.Writer
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil, output: &bytes.Buffer{}}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	env.output = outer.output
	return env
}

func (e *Environment) Get(key string) (Object, bool) {
	obj, ok := e.store[key]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(key)
	}
	return obj, ok
}

func (e *Environment) Set(key string, val Object) Object {
	e.store[key] = val
	return val
}

func (e *Environment) SetOutput(out io.Writer) {
	e.output = out
}

func (e *Environment) GetOutput() io.Writer {
	return e.output
}
