package pgsql_test

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/samlitowitz/expressions"
	"github.com/samlitowitz/expressions/pgsql"
)

func TestWhereClauseFromExpression(t *testing.T) {
	dburl, err := pgsqlDBURLFromEnv()
	if err != nil {
		t.Fatal("pgsql: dburl: ", err)
	}
	db, err := sql.Open("pgx", dburl)
	if err != nil {
		t.Fatal("pgsql: ", err)
	}
	t.Cleanup(func() {
		err := db.Close()
		if err != nil {
			t.Fatal("pgsql: ", err)
		}
	})
	_, err = db.Exec(`
DROP TABLE IF EXISTS "test";
CREATE TABLE IF NOT EXISTS "test" (
    "id" INTEGER,
    "data" TEXT,
    "timestamp" TEXT /* stored as RFC3339 string */,

    PRIMARY KEY (
        "id"
    )
);

INSERT INTO "test" ("id", "data", "timestamp") VALUES (1, 'one', '2000-01-01 00:00:00');
INSERT INTO "test" ("id", "data", "timestamp") VALUES (2, 'two', '2000-01-02 00:00:00');
INSERT INTO "test" ("id", "data", "timestamp") VALUES (3, 'three', '2000-01-03 00:00:00');
INSERT INTO "test" ("id", "data", "timestamp") VALUES (4, 'four', '2000-01-04 00:00:00');
INSERT INTO "test" ("id", "data", "timestamp") VALUES (5, 'five', '2000-01-05 00:00:00');
`)
	if err != nil {
		t.Fatal("sqlite: ", err)
	}

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
		clauses, binds, err := pgsql.WhereClauseFromExpression(testCase.expression, 1, idMap)
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

func pgsqlDBURLFromEnv() (string, error) {
	host := os.Getenv("DB_HOST")
	if len(host) == 0 {
		return "", fmt.Errorf("no host provided")
	}
	userFile := os.Getenv("DB_USER_FILE")
	if len(userFile) == 0 {
		return "", fmt.Errorf("no user file provided")
	}
	passwordFile := os.Getenv("DB_PASSWORD_FILE")
	if len(passwordFile) == 0 {
		return "", fmt.Errorf("no password file provided")
	}

	user, err := os.ReadFile(userFile)
	if err != nil {
		return "", fmt.Errorf("pgsql dburl from env: user: %w", err)
	}
	password, err := os.ReadFile(passwordFile)
	if err != nil {
		return "", fmt.Errorf("pgsql dburl from env: password: %w", err)
	}
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s/%s",
		strings.TrimSpace(string(user)),
		strings.TrimSpace(string(password)),
		host,
		strings.TrimSpace(string(user)),
	)

	connConfig, err := pgx.ParseConfig(dbURL)
	if err != nil {
		return "", fmt.Errorf("parse config: %w", err)
	}
	connStr := stdlib.RegisterConnConfig(connConfig)
	return connStr, nil
}
