package expressions_test

import (
	"fmt"
	"github.com/samlitowitz/expressions"
)

func ExampleAnd_aAndB() {
	fmt.Println(
		expressions.NewAnd(
			expressions.NewIdentifier("A"),
			expressions.NewIdentifier("B"),
		),
	)
	// Output: A AND B
}

func ExampleOr_aOrB() {
	fmt.Println(
		expressions.NewOr(
			expressions.NewIdentifier("A"),
			expressions.NewIdentifier("B"),
		),
	)
	// Output: A OR B
}

func ExampleNot_notA() {
	fmt.Println(
		expressions.NewNot(
			expressions.NewIdentifier("A"),
		),
	)
	// Output: NOT A
}
