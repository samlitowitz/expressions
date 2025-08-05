package pgsql

import (
	"fmt"
	"time"

	"github.com/samlitowitz/expressions"
)

// WhereClauseFromExpression returns a string representation of the clauses, the associated binds, and any errors
// that may have occurred while building the these.
func WhereClauseFromExpression(expr expressions.Expression, paramIdx int, idToSQLiteIdent map[expressions.ID]string) (string, []any, error) {
	if expr == nil {
		return "", nil, nil
	}
	switch expr := expr.(type) {
	case *expressions.And:
		left, leftBinds, err := WhereClauseFromExpression(expr.Left(), paramIdx, idToSQLiteIdent)
		if err != nil {
			return "", nil, err
		}
		right, rightBinds, err := WhereClauseFromExpression(expr.Right(), paramIdx+len(leftBinds), idToSQLiteIdent)
		if err != nil {
			return "", nil, err
		}
		return fmt.Sprintf("%s AND %s", left, right), append(leftBinds, rightBinds...), nil

	case *expressions.Or:
		left, leftBinds, err := WhereClauseFromExpression(expr.Left(), paramIdx, idToSQLiteIdent)
		if err != nil {
			return "", nil, err
		}
		right, rightBinds, err := WhereClauseFromExpression(expr.Right(), paramIdx+len(leftBinds), idToSQLiteIdent)
		if err != nil {
			return "", nil, err
		}
		return fmt.Sprintf("%s OR %s", left, right), append(leftBinds, rightBinds...), nil
	case *expressions.Not:
		operand, binds, err := WhereClauseFromExpression(expr.Operand(), paramIdx, idToSQLiteIdent)
		if err != nil {
			return "", nil, err
		}
		return fmt.Sprintf("NOT %s", operand), binds, nil

	case *expressions.Equals:
		left, leftBinds, err := WhereClauseFromExpression(expr.Left(), paramIdx, idToSQLiteIdent)
		if err != nil {
			return "", nil, err
		}
		right, rightBinds, err := WhereClauseFromExpression(expr.Right(), paramIdx+len(leftBinds), idToSQLiteIdent)
		if err != nil {
			return "", nil, err
		}
		return fmt.Sprintf("%s = %s", left, right), append(leftBinds, rightBinds...), nil

	case *expressions.Identifier:
		colName, ok := idToSQLiteIdent[expr.ID()]
		if !ok {
			return "", nil, fmt.Errorf("unmapped ID: %s", expr.ID())
		}
		return colName, nil, nil
	case *expressions.Scalar:
		return fmt.Sprintf("$%d", paramIdx), []any{expr.Value()}, nil
	case expressions.Timestamp:
		return fmt.Sprintf("$%d", paramIdx), []any{time.Time(expr)}, nil
	default:
		return "", nil, fmt.Errorf("unknown expression")
	}
}
