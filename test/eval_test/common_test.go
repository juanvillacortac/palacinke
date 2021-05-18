package eval_test

import (
	"testing"

	"github.com/juanvillacortac/palacinke/pkg/eval"
	"github.com/juanvillacortac/palacinke/pkg/lexer"
	"github.com/juanvillacortac/palacinke/pkg/object"
	"github.com/juanvillacortac/palacinke/pkg/parser"
)

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	return eval.Eval(program, env)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object isn't Integer. Got: %T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. Got: %d, want: %d",
			result.Value, expected)
		return false
	}
	return true
}

func testBooleanObject(t *testing.T, obj object.Object, expected bool) bool {
	result, ok := obj.(*object.Boolean)
	if !ok {
		t.Errorf("object isn't Boolean. Got: %T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. Got: %t, want: %t",
			result.Value, expected)
		return false
	}
	return true
}

func testNilObject(t *testing.T, obj object.Object) bool {
	if obj != eval.NIL {
		t.Errorf("object is not NIL. Got: %T (%+v)", obj, obj)
		return false
	}
	return true
}
