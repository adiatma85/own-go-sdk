package sql

import (
	"context"
	"database/sql"

	"github.com/adiatma85/own-go-sdk/instrument"
	"github.com/jmoiron/sqlx"
)

type CommandStmt interface {
	Close() error
	Select(name string, dest interface{}, args ...interface{}) error
	Get(name string, dest interface{}, args ...interface{}) error
	QueryRow(name string, args ...interface{}) (*sqlx.Row, error)
	Query(name string, args ...interface{}) (*sqlx.Rows, error)
	Exec(name string, args ...interface{}) (sql.Result, error)
}

type commandStmt struct {
	ctx           context.Context
	name          string
	stmt          *sqlx.Stmt
	instrument    instrument.Interface
	useInstrument bool
}

func initStmt(ctx context.Context, name string, stmt *sqlx.Stmt, instr instrument.Interface, useInstr bool) CommandStmt {
	return &commandStmt{
		ctx:           ctx,
		name:          name,
		stmt:          stmt,
		instrument:    instr,
		useInstrument: useInstr,
	}
}

func (x *commandStmt) Close() error {
	return x.stmt.Close()
}

func (x *commandStmt) Select(name string, dest interface{}, args ...interface{}) error {
	if x.useInstrument {
		timer := x.instrument.MySQLQueryTimer(name)
		defer timer.ObserveDuration()
	}
	return x.stmt.SelectContext(x.ctx, dest, args...)
}

func (x *commandStmt) Get(name string, dest interface{}, args ...interface{}) error {
	if x.useInstrument {
		timer := x.instrument.MySQLQueryTimer(name)
		defer timer.ObserveDuration()
	}
	return x.stmt.GetContext(x.ctx, dest, args...)
}

func (x *commandStmt) QueryRow(name string, args ...interface{}) (*sqlx.Row, error) {
	if x.useInstrument {
		timer := x.instrument.MySQLQueryTimer(name)
		defer timer.ObserveDuration()
	}
	return x.stmt.QueryRowxContext(x.ctx, args...), nil
}

func (x *commandStmt) Query(name string, args ...interface{}) (*sqlx.Rows, error) {
	if x.useInstrument {
		timer := x.instrument.MySQLQueryTimer(name)
		defer timer.ObserveDuration()
	}
	return x.stmt.QueryxContext(x.ctx, args...)
}

func (x *commandStmt) Exec(name string, args ...interface{}) (sql.Result, error) {
	if x.useInstrument {
		timer := x.instrument.MySQLQueryTimer(name)
		defer timer.ObserveDuration()
	}
	return x.stmt.ExecContext(x.ctx, args...)
}
