package pgscanbench_test

import (
	"context"
	"testing"
	"time"

	scany "github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	randallmlough "github.com/randallmlough/pgxscan"
)

const defaultDbURI = "postgres://user:password@localhost:5432/db?sslmode=disable"

type Name struct {
	Name struct {
		Firstname string
		Lastname  string
	}
}

type User struct {
	Id    int64
	Title string
	Body  string
	Tt    time.Time
	Count int
	Jj    Name
}

func BenchmarkRandallmlough(b *testing.B) {
	testDB, err := pgxpool.Connect(context.Background(), defaultDbURI)
	if err != nil {
		b.Fail()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rows, _ := testDB.Query(context.Background(), `SELECT * FROM "bench"`)
		var users []User
		if err := randallmlough.NewScanner(rows).Scan(&users); err != nil {
			b.Errorf("BenchmarkRandallmlough() failed to scan into []User. Reason:  %v", err)
			b.FailNow()
		}
		_ = len(users)
	}
}

func BenchmarkScany(b *testing.B) {
	testDB, err := pgxpool.Connect(context.Background(), defaultDbURI)
	if err != nil {
		b.Fail()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rows, _ := testDB.Query(context.Background(), `SELECT * FROM "bench"`)
		var users []User
		if err := scany.ScanAll(&users, rows); err != nil {
			b.Errorf("BenchmarkScany() failed to scan into []User. Reason:  %v", err)
			b.FailNow()
		}
		_ = len(users)
	}
}
