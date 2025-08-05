package expressions_test

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/samlitowitz/expressions"
)

func ExampleIdentifier_arbitrarilyMappedIdentifier() {
	fmt.Println(expressions.NewIdentifier(expressions.ID(uuid.New().String())))
}

func ExampleIdentifier_structField() {
	fmt.Println(expressions.NewIdentifier("struct.field"))
	// Output: struct.field
}

func ExampleIdentifier_tableColumn() {
	fmt.Println(expressions.NewIdentifier("table.column"))
	// Output: table.column
}
