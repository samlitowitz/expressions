package expressions

import "fmt"

var _ Expression = (*And)(nil)
var _ Expression = (*Or)(nil)
var _ Expression = (*Not)(nil)

// And represents a logical and expression.
type And struct {
	*Binary
}

// NewAnd creates a new [And] expression.
func NewAnd(left, right Expression) *And {
	return &And{
		&Binary{
			left:  left,
			right: right,
		},
	}
}

func (a And) String() string {
	return fmt.Sprintf("%s AND %s", a.left, a.right)
}

// Or represents a logical or expression.
type Or struct {
	*Binary
}

// NewOr creates a new [Or] expression.
func NewOr(left, right Expression) *Or {
	return &Or{
		&Binary{
			left:  left,
			right: right,
		},
	}
}

func (o Or) String() string {
	return fmt.Sprintf("%s OR %s", o.left, o.right)
}

// Not represents a logical not expression.
type Not struct {
	expr Expression
}

// NewNot creates a new [Not] expression.
func NewNot(expr Expression) *Not {
	return &Not{
		expr: expr,
	}
}

func (expr Not) Operands() []Expression {
	return []Expression{expr.expr}
}

func (expr Not) Operand() Expression {
	return expr.expr
}

func (n Not) String() string {
	return fmt.Sprintf("NOT %s", n.expr)
}
