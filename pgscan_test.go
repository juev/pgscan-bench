package pgscanbench_test

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
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

// Make the Attrs struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a Name) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Make the Attrs struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *Name) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
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

func BenchmarkRandallmloughScanOne(b *testing.B) {
	testDB, err := pgxpool.Connect(context.Background(), defaultDbURI)
	if err != nil {
		b.FailNow()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rows, _ := testDB.Query(context.Background(), `SELECT * FROM "bench" LIMIT 1`)
		var user User
		if err := randallmlough.NewScanner(rows).Scan(&user); err != nil {
			b.Errorf("BenchmarkRandallmloughScanOne() failed to scan into []User. Reason:  %v", err)
			b.FailNow()
		}
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

func BenchmarkScanyScanOne(b *testing.B) {
	testDB, err := pgxpool.Connect(context.Background(), defaultDbURI)
	if err != nil {
		b.FailNow()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rows, _ := testDB.Query(context.Background(), `SELECT * FROM "bench" LIMIT 1`)
		var user User
		if err := scany.ScanOne(&user, rows); err != nil {
			b.Errorf("BenchmarkScanyScanOne() failed to scan into []User. Reason:  %v", err)
			b.FailNow()
		}
	}
}

func BenchmarkManual(b *testing.B) {
	testDB, err := pgxpool.Connect(context.Background(), defaultDbURI)
	if err != nil {
		b.FailNow()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rows, _ := testDB.Query(context.Background(), `SELECT * FROM "bench"`)
		var users []User
		for rows.Next() {
			var user User
			if err := rows.Scan(&user.Id, &user.Title, &user.Body, &user.Tt, &user.Count, &user.Jj); err != nil {
				b.FailNow()
			}
			users = append(users, user)
		}
		_ = len(users)
	}
}

func BenchmarkManualScanOne(b *testing.B) {
	testDB, err := pgxpool.Connect(context.Background(), defaultDbURI)
	if err != nil {
		b.FailNow()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		row := testDB.QueryRow(context.Background(), `SELECT * FROM "bench" LIMIT 1`)
		var user User
		if err := row.Scan(&user.Id, &user.Title, &user.Body, &user.Tt, &user.Count, &user.Jj); err != nil {
			b.FailNow()
		}
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

func ExampleManual() {
	testDB, _ := pgxpool.Connect(context.Background(), defaultDbURI)
	rows, _ := testDB.Query(context.Background(), `SELECT * FROM "bench"`)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.Title, &user.Body, &user.Tt, &user.Count, &user.Jj); err != nil {
			fmt.Printf("error: %v\n", err)
		}
		fmt.Printf("%#v", user)
		break
	}
	// Output:
	// pgscanbench_test.User{Id:1, Title:"aee393a97def6cc5f1e1811a2d079e7c", Body:"81db91c31b4202372436050951288536", Tt:time.Date(1949, time.July, 16, 23, 13, 44, 709453000, time.UTC), Count:41, Jj:pgscanbench_test.Name{Name:struct { Firstname string; Lastname string }{Firstname:"first", Lastname:"first2"}}}
}
