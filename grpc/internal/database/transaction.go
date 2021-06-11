package database

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

type Tx struct {
	*sqlx.Tx
	DB *DB
}

// A Txfn is a function that will be called with an initialized `Transaction` object
// that can be used for executing statements and queries against a database.
type TxFn func(context.Context, Tx) error

// WithTransaction creates a new transaction and handles rollback/commit based on the
// error object returned by the `TxFn`
func (db *DB) WithTransaction(ctx context.Context, fn TxFn) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "db.WithTransaction")
	defer span.Finish()

	sqlTx, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	tx := Tx{Tx: sqlTx, DB: db}

	defer func() {
		if p := recover(); p != nil {
			// a panic occurred, rollback and repanic
			span.SetTag("error", true)
			tx.Rollback()
			panic(p)
		} else if err != nil {
			// something went wrong, rollback
			span.SetTag("error", true)
			tx.Rollback()
		} else {
			// all good, commit
			err = tx.Commit()
		}
	}()

	err = fn(ctx, tx)
	return err
}

// SelectContext calls SelectContext and trace the function call.
func (tx *Tx) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, tx.DB.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", args)

	return tx.Tx.SelectContext(ctx, dest, query, args...)
}

// GetContext calls GetContext and trace the function call.
func (tx *Tx) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, tx.DB.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", args)

	return tx.Tx.GetContext(ctx, dest, query, args...)
}

// ExecContext calls ExecContext and trace the function call.
func (tx *Tx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, tx.DB.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", args)

	return tx.Tx.ExecContext(ctx, query, args...)
}

// QueryContext calls QueryContext and trace the function call.
func (tx *Tx) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, tx.DB.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", args)

	return tx.Tx.QueryContext(ctx, query, args...)
}

// QueryRowContext calls QueryContext and trace the function call.
func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, tx.DB.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", args)

	return tx.Tx.QueryRowContext(ctx, query, args...)
}
