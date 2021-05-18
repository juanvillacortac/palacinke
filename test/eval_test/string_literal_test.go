package eval_test

import (
	"testing"

	"github.com/juanvillacortac/palacinke/pkg/object"
)

func TestStringLiteral(t *testing.T) {
	input := `"Hello World!"`

	evaluated := testEval(input)
	str, ok := evaluated.(*object.String)
	if !ok {
		t.Fatalf("object is not String. Got: %T (%+v)", evaluated, evaluated)
	}

	if str.Value != "Hello World!" {
		t.Errorf("String has wrong value. Got: %q", str.Value)
	}
}
