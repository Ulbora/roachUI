package database

import (
	crdb "github.com/GolangToolKits/go-cockroachdb"
)

type Database interface {
	Create(query string) bool
	Drop(query string) bool
	Insert(table string, params map[string]any) (bool, int64)
	Update(table string, params map[string]any) bool
	Get(table string, params map[string]any) *crdb.DbRow
	GetList(table string, params map[string]any) *crdb.DbRows
	Delete(table string, params map[string]any) bool
}
