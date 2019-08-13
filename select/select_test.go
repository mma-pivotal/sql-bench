package main

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const dblink = //please setup your connection string here, must be DSN format

func selectRecord(b *testing.B, db *sql.DB) {
	_, err := db.Exec("select * from test_table where id=10001;")
	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkSelectMaxOpenConns1(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxOpenConns(1)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			selectRecord(b, db)
		}
	})
}

func BenchmarkSelectMaxOpenConns2(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxOpenConns(2)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			selectRecord(b, db)
		}
	})
}

func BenchmarkSelectMaxOpenConns4(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxOpenConns(4)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			selectRecord(b, db)
		}
	})
}

func BenchmarkSelectMaxOpenConns8(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxOpenConns(8)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			selectRecord(b, db)
		}
	})
}

func BenchmarkSelectMaxOpenConns16(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxOpenConns(16)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			selectRecord(b, db)
		}
	})
}

func BenchmarkSelectMaxOpenConns30(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxOpenConns(30)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			selectRecord(b, db)
		}
	})
}

func BenchmarkSelectMaxOpenConnsUnlimited(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			selectRecord(b, db)
		}
	})
}

func BenchmarkMaxIdleConnsNone(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxIdleConns(0)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			selectRecord(b, db)
		}
	})
}

func BenchmarkMaxIdleConns1(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxIdleConns(1)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			selectRecord(b, db)
		}
	})
}

func BenchmarkMaxIdleConns2(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxIdleConns(2)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			selectRecord(b, db)
		}
	})
}

func BenchmarkMaxIdleConns4(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxIdleConns(4)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			selectRecord(b, db)
		}
	})
}

func BenchmarkMaxIdleConns8(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxIdleConns(8)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			selectRecord(b, db)
		}
	})
}
func BenchmarkMaxIdleConns16(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxIdleConns(16)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			selectRecord(b, db)
		}
	})
}

func BenchmarkMaxIdleConns30(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxIdleConns(30)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			selectRecord(b, db)
		}
	})
}
