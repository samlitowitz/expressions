package expressions

import "fmt"

var _ Expression = (*Binary)(nil)

// Expression is a unified representation of any expression.
type Expression interface {
	// Operands returns a list of the expression's operands.
	Operands() []Expression
	// String returns an implementation agnostic string representation of the expression.
	String() string
}

// Binary represents all binary expressions.
//
// This structure can be embedded in another to represent more specific
// binary expressions. For examples, see expressions in this package
// which do so such as [And], [Equals], [Identifier], and [Scalar].
type Binary struct {
	left  Expression
	right Expression
}

// NewBinary creates a new [Binary] expression.
func NewBinary(left Expression, right Expression) *Binary {
	return &Binary{left: left, right: right}
}

func (expr Binary) Operands() []Expression {
	return []Expression{expr.left, expr.right}
}

// Left returns the left operand.
func (expr Binary) Left() Expression {
	return expr.left
}

// Right returns the right operand.
func (expr Binary) Right() Expression {
	return expr.right
}

func (expr Binary) String() string {
	return fmt.Sprintf("%s ?? %s", expr.left, expr.right)
}
