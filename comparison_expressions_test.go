package expressions_test

import (
	"fmt"
	"github.com/samlitowitz/expressions"
)

func ExampleEquals_aEqualsB() {
	fmt.Println(
		expressions.NewEquals(
			expressions.NewIdentifier("A"),
			expressions.NewIdentifier("B"),
		),
	)
	// Output: A = B
}
