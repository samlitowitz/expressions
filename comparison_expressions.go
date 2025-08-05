package expressions

import "fmt"

var _ Expression = (*Equals)(nil)

// Equals represents an expression of equality.
type Equals struct {
	*Binary
}

// NewEquals creates a new [Equals] expression.
func NewEquals(left, right Expression) *Equals {
	return &Equals{
		&Binary{
			left:  left,
			right: right,
		},
	}
}

func (e Equals) String() string {
	return fmt.Sprintf("%s = %s", e.left, e.right)
}
