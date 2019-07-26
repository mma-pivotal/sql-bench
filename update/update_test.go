package main

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const dblink = "sqlbench:Password01.@tcp(10.193.76.17:3306)/mysql" //please setup your connection string here, must be DSN format

func updateRecord(b *testing.B, db *sql.DB) {
	_, err := db.Exec("update test_table set name='1st' , description='1st record' where id=1;")
	if err != nil {
		b.Fatal(err)
	}
}

func BenchmarkMaxOpenConns1(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxOpenConns(1)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			updateRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConns2(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxOpenConns(2)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			updateRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConns4(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxOpenConns(4)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			updateRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConns8(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxOpenConns(8)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			updateRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConns16(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxOpenConns(16)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			updateRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConns30(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxOpenConns(30)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			updateRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConnsUnlimited(b *testing.B) {
	db, err := sql.Open("mysql", dblink)
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			updateRecord(b, db)
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
			updateRecord(b, db)
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
			updateRecord(b, db)
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
			updateRecord(b, db)
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
			updateRecord(b, db)
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
			updateRecord(b, db)
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
			updateRecord(b, db)
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
			updateRecord(b, db)
		}
	})
}
