package expressions

import (
	"fmt"
	"time"
)

var _ Expression = (*Scalar)(nil)
var _ Expression = (*Timestamp)(nil)

// Scalar represents a scalar value used in an expression.
type Scalar struct {
	value any
}

// NewScalar creates a new [Scalar] expression.
func NewScalar(value any) *Scalar {
	return &Scalar{
		value: value,
	}
}

func (s Scalar) Operands() []Expression {
	return nil
}

// Value returns the scalar value represented by [Scalar].
func (s Scalar) Value() any {
	return s.value
}

func (s Scalar) String() string {
	return fmt.Sprintf("%#v", s.value)
}

// Timestamp represents a timestamp value used in an expression.
type Timestamp time.Time

// NewTimestamp creates a new [Timestamp] expression.
func NewTimestamp(value time.Time) Timestamp {
	return Timestamp(value)
}

func (t Timestamp) Operands() []Expression {
	return nil
}

func (t Timestamp) String() string {
	return fmt.Sprintf("%v", time.Time(t).Format(time.RFC3339))
}
