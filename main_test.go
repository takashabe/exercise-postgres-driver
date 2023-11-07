package main

import (
	"context"
	"os"
	"testing"

	"github.com/k0kubun/pp/v3"
	"gopkg.in/testfixtures.v2"
)

func TestMain(m *testing.M) {
	if err := loadFixtures(); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func loadFixtures() error {
	testfixtures.SkipDatabaseNameCheck(true)
	testCtx, err := testfixtures.NewFolder(
		dbMap.Db,
		&testfixtures.PostgreSQL{
			SkipResetSequences: true,
		},
		"fixtures",
	)
	if err != nil {
		return err
	}
	if err := testCtx.Load(); err != nil {
		return err
	}
	pp.Println("fixtures loaded")
	return nil
}

func TestReadPersons(t *testing.T) {
	ctx := context.Background()

	persons, err := readPersons(ctx)
	if err != nil {
		t.Fatal(err)
	}
	if len(persons) != 6 {
		t.Fatalf("expected 2 persons, got %d", len(persons))
	}
}
