//go:build !no_libsql
// +build !no_libsql

package main

import (
	_ "github.com/pressly/goose/v3/lib/github.com/tursodatabase/go-libsql"
)
