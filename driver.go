package main

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"

	"github.com/k0kubun/pp/v3"
	"github.com/lib/pq"
)

type tenantConnector struct {
	dsn      string
	tenantID int64
	conn     driver.Conn
}

// Begin implements driver.Conn.
func (c *tenantConnector) Begin() (driver.Tx, error) {
	return c.conn.Begin()
}

// Close implements driver.Conn.
func (c *tenantConnector) Close() error {
	return c.conn.Close()
}

// Prepare implements driver.Conn.
func (c *tenantConnector) Prepare(query string) (driver.Stmt, error) {
	return c.conn.Prepare(query)
}

func (c *tenantConnector) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	return c.conn.(driver.ExecerContext).ExecContext(ctx, query, args)
}

func (tc *tenantConnector) Connect(ctx context.Context) (driver.Conn, error) {
	conn, err := pq.NewConnector(tc.dsn)
	if err != nil {
		return nil, err
	}
	dbConn, err := conn.Connect(ctx)
	if err != nil {
		return nil, err
	}
	tc.conn = dbConn
	pp.Println("setting tenant_id to", tc.tenantID)
	if execer, ok := tc.conn.(driver.ExecerContext); ok {
		_, err = execer.ExecContext(ctx, fmt.Sprintf("SET app.tenant_id = %d", tc.tenantID), nil)
		if err != nil {
			dbConn.Close()
			return nil, err
		}
	} else {
		// Handle the case where ExecContext is not supported
		return nil, errors.New("connection does not support ExecContext")
	}
	return tc, nil
}

func (tc *tenantConnector) Driver() driver.Driver {
	return &tenantDriver{}
}

func (tc *tenantConnector) ResetSession(ctx context.Context) error {
	pp.Println("resetting tenant_id")
	if execer, ok := tc.conn.(driver.ExecerContext); ok {
		_, err := execer.ExecContext(ctx, "RESET ROLE;", nil)
		return err
	}
	return errors.New("connection does not support ExecContext")
}

type tenantDriver struct{}

func (d *tenantDriver) Open(name string) (driver.Conn, error) {
	// Should not be used directly.
	return nil, errors.New("tenantDriver does not support the Open method")
}
