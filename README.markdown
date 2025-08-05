# Expressions

[![Go Report Card][goreportcardbadge]][goreportcard]
[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)][godev]

This package is intended to allow the usage of implementation agnostic expressions with a focus on expressions relevant
to SQL.

This package contains functionality to use the provided expressions with [PgSQL][pgsqldoc]
and [SQLite][sqlitedoc].

See the [documentation][godev] for more information.

[goreportcardbadge]: https://goreportcard.com/badge/github.com/samlitowitz/expressions
[goreportcard]: https://goreportcard.com/report/github.com/samlitowitz/expressions

[godev]: https://pkg.go.dev/github.com/samlitowitz/expressions

[pgsqldoc]: https://pkg.go.dev/github.com/samlitowitz/expressions/pgsql#WhereClauseFromExpression

[sqlitedoc]: https://pkg.go.dev/github.com/samlitowitz/expressions/sqlite#WhereClauseFromExpression

# Install

```
go get -u https://github.com/samlitowitz/expressions
```

# License

Apache 2.0 - See [LICENSE][license]

[license]: https://github.com/samlitowitz/expressions/blob/master/LICENSE
