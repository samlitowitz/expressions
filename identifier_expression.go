package expressions

import "fmt"

var _ Expression = (*Identifier)(nil)

// The ID type is a string representing an identifier to be used by the [Identifier] expression.
type ID string

// Identifier represents an identifier used in an expression.
// Examples of identifiers are SQL table column names and struct field names.
type Identifier struct {
	id ID
}

// NewIdentifier creates a new [Identifier].
func NewIdentifier(id ID) *Identifier {
	return &Identifier{
		id: id,
	}
}

func (expr Identifier) Operands() []Expression {
	return nil
}

// ID returns the string representation of the [Identifier].
func (expr Identifier) ID() ID {
	return expr.id
}

func (expr Identifier) String() string {
	return fmt.Sprintf("%s", expr.id)
}
