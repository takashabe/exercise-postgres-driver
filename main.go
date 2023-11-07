package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/k0kubun/pp/v3"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v2"
)

var (
	db    *sql.DB
	dbMap *gorp.DbMap
)

type Person struct {
	ID        int64     `db:"id"`
	TenantID  int64     `db:"tenant_id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	BirthDate time.Time `db:"birth_date"`
	CreatedAt time.Time `db:"created_at"`
}

func init() {
	db, err := openDB(context.TODO())
	if err != nil {
		panic(err)
	}
	dbMap = &gorp.DbMap{
		Db:      db,
		Dialect: gorp.PostgresDialect{},
	}
}

func main() {
	ctx := context.Background()

	persons, err := readPersons(ctx)
	if err != nil {
		panic(err)
	}
	pp.Println(persons)
}

func readPersons(ctx context.Context) ([]Person, error) {
	persons := []Person{}
	if _, err := dbMap.Select(&persons, `select * from persons`); err != nil {
		return nil, err
	}
	return persons, nil
}

func openDB(ctx context.Context) (*sql.DB, error) {
	dsn := "postgres://postgres:password@localhost:5434/postgres?sslmode=disable"
	// dsn := "postgres://app:password@localhost:5434/postgres?sslmode=disable"
	db := sql.OpenDB(&tenantConnector{
		dsn:      dsn,
		tenantID: 1,
	})
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	return db, nil
}
