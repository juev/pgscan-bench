package pgscanbench_test

import (
	"testing"

	_ "github.com/georgysavva/scany/pgxscan"
	_ "github.com/randallmlough/pgxscan"
)

func BenchmarkScany(b *testing.B) {
	//
}
