package sqlite_test

import (
	"database/sql"
	"testing"

	"github.com/samlitowitz/expressions"
	"github.com/samlitowitz/expressions/sqlite"

	_ "modernc.org/sqlite"
)

func TestWhereClauseFromExpression(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatal("sqlite: ", err)
	}
	t.Cleanup(func() {
		err := db.Close()
		if err != nil {
			t.Fatal("sqlite: ", err)
		}
	})
	stmt, err := db.Prepare(`
DROP TABLE IF EXISTS "test";
CREATE TABLE IF NOT EXISTS "test" (
    "id" INTEGER,
    "data" TEXT,
    "timestamp" TEXT /* stored as RFC3339 string */,

    PRIMARY KEY (
        "id"
    )
);

INSERT INTO "test" ("id", "data", "timestamp") VALUES (1, "one", "2000-01-01 00:00:00");
INSERT INTO "test" ("id", "data", "timestamp") VALUES (2, "two", "2000-01-02 00:00:00");
INSERT INTO "test" ("id", "data", "timestamp") VALUES (3, "three", "2000-01-03 00:00:00");
INSERT INTO "test" ("id", "data", "timestamp") VALUES (4, "four", "2000-01-04 00:00:00");
INSERT INTO "test" ("id", "data", "timestamp") VALUES (5, "five", "2000-01-05 00:00:00");
`)
	if err != nil {
		t.Fatal("sqlite: ", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		t.Fatal("sqlite: ", err)
	}
	t.Cleanup(func() {
		err := stmt.Close()
		if err != nil {
			t.Fatal("sqlite: ", err)
		}
	})

	idMap := map[expressions.ID]string{
		`"test"."id"`:        `"test"."id"`,
		`"test"."timestamp"`: `"test"."timestamp"`,
	}

	testCases := map[string]struct {
		expression expressions.Expression
		expectedID int
	}{
		"and": {
			expression: expressions.NewAnd(
				expressions.NewEquals(
					expressions.NewIdentifier(`"test"."id"`),
					expressions.NewScalar(1),
				),
				expressions.NewEquals(
					expressions.NewIdentifier(`"test"."timestamp"`),
					expressions.NewScalar("2000-01-01 00:00:00"),
				),
			),
			expectedID: 1,
		},
		"or": {
			expression: expressions.NewOr(
				expressions.NewEquals(
					expressions.NewIdentifier(`"test"."id"`),
					expressions.NewScalar(1),
				),
				expressions.NewEquals(
					expressions.NewIdentifier(`"test"."timestamp"`),
					expressions.NewScalar("2000-01-01 00:00:00"),
				),
			),
			expectedID: 1,
		},
		"not": {
			expression: expressions.NewAnd(
				expressions.NewNot(
					expressions.NewEquals(
						expressions.NewIdentifier(`"test"."id"`),
						expressions.NewScalar(2),
					),
				),
				expressions.NewAnd(
					expressions.NewNot(
						expressions.NewEquals(
							expressions.NewIdentifier(`"test"."id"`),
							expressions.NewScalar(3),
						),
					),
					expressions.NewAnd(
						expressions.NewNot(
							expressions.NewEquals(
								expressions.NewIdentifier(`"test"."id"`),
								expressions.NewScalar(4),
							),
						),
						expressions.NewNot(
							expressions.NewEquals(
								expressions.NewIdentifier(`"test"."id"`),
								expressions.NewScalar(5),
							),
						),
					),
				),
			),
			expectedID: 1,
		},
	}
	for testDesc, testCase := range testCases {
		clauses, binds, err := sqlite.WhereClauseFromExpression(testCase.expression, idMap)
		if err != nil {
			t.Fatal(testDesc, ": ", err)
		}
		query := `SELECT "test"."id" FROM "test" WHERE ` + clauses + " LIMIT 1"
		stmt, err := db.Prepare(query)
		if err != nil {
			t.Fatal(testDesc, ": ", err)
		}
		t.Cleanup(func() {
			_ = stmt.Close()
		})
		rows, err := stmt.Query(binds...)
		if err != nil {
			t.Fatal(testDesc, ": ", err)
		}
		t.Cleanup(func() {
			_ = rows.Close()
		})
		for rows.Next() {
			actualID := 0

			if err = rows.Scan(&actualID); err != nil {
				t.Fatal(testDesc, ": ", err)
			}
		}
		if err = rows.Err(); err != nil {
			t.Fatal(testDesc, ": ", err)
		}
		err = rows.Close()
		if err != nil {
			t.Fatal(testDesc, ": ", err)
		}
		err = stmt.Close()
		if err != nil {
			t.Fatal(testDesc, ": ", err)
		}
	}
}
