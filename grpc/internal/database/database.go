package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

//Driver list
const (
	DriverPostgres = "postgres"
)

const SpanName = "db.query"

type (
	DB struct {
		*sqlx.DB
	}

	//DBConfig for databases configuration
	DBConfig struct {
		PrimaryDSN      string
		ReplicaDSN      string
		MaxIdleConn     int
		MaxConn         int
		ConnMaxLifetime time.Duration
	}

	//Store is used to persist DB connection
	Store struct {
		Primary *DB
		Replica *DB
	}
)

func (s *Store) GetPrimary() *DB {
	return s.Primary
}

func (s *Store) GetReplica() *DB {
	return s.Replica
}

func New(cfg DBConfig, dbDriver string) *Store {
	primary := sqlx.MustConnect(dbDriver, cfg.PrimaryDSN)
	primary.SetMaxOpenConns(cfg.MaxConn)
	primary.SetMaxIdleConns(cfg.MaxIdleConn)
	primary.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	replica := sqlx.MustConnect(dbDriver, cfg.ReplicaDSN)
	replica.SetMaxOpenConns(cfg.MaxConn)
	replica.SetMaxIdleConns(cfg.MaxIdleConn)
	replica.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	return &Store{Primary: &DB{DB: primary}, Replica: &DB{DB: replica}}
}

//GetContext calls GetContext and trace the function call.
func (db *DB) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, db.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", args)

	return db.DB.GetContext(ctx, dest, query, args...)
}

//SelectContext calls SelectContext and trace the function call.
func (db *DB) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, db.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", args)

	return db.DB.SelectContext(ctx, dest, query, args...)
}

//ExecContext calls ExecContext and trace the function call.
func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, db.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", args)

	return db.DB.ExecContext(ctx, query, args...)
}

//QueryContext calls QueryContext and trace the function call.
func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, db.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", args)

	return db.DB.QueryContext(ctx, query, args...)
}

// QueryRowContext calls QueryRowContext and trace the function call.
func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	span, ctx := opentracing.StartSpanFromContext(ctx, SpanName)
	defer span.Finish()

	ext.DBStatement.Set(span, query)
	ext.DBInstance.Set(span, db.DriverName())
	ext.DBType.Set(span, "sql")
	span.SetTag("db.values", args)

	return db.DB.QueryRowContext(ctx, query, args...)
}
