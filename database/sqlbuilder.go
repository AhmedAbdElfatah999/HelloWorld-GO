package database

import (
	"fmt"
	"strings"
)

type SqlBuilder struct {
}

func NewSqlBuilder() *SqlBuilder {
	return &SqlBuilder{}
}

func (d *SqlBuilder) Insert(tableName string, argsKeys []string, argsVals []string) string {
	sql := fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)", tableName, strings.Join(argsKeys, ","), strings.Join(argsVals, ","))
	return sql
}

func (d *SqlBuilder) SelectAll(tableName string) string {
	sql := fmt.Sprintf("SELECT * FROM %s", tableName)
	return sql
}