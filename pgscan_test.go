package pgscanbench_test

import (
	"context"
	"fmt"
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
		b.FailNow()
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
		b.FailNow()
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

func ExampleScany() {
	var users []User
	testDB, _ := pgxpool.Connect(context.Background(), defaultDbURI)
	rows, _ := testDB.Query(context.Background(), `SELECT * FROM "bench"`)
	if err := scany.ScanAll(&users, rows); err == nil {
		fmt.Printf("%#v", users[0])
	}
	// Output:
	// pgscanbench_test.User{Id:1, Title:"aee393a97def6cc5f1e1811a2d079e7c", Body:"81db91c31b4202372436050951288536", Tt:time.Date(1949, time.July, 16, 23, 13, 44, 709453000, time.UTC), Count:41, Jj:pgscanbench_test.Name{Name:struct { Firstname string; Lastname string }{Firstname:"first", Lastname:"first2"}}}
}
