// File: test/helpers_test.go
package test

import (
	"crypto/rsa"
	"crypto/x509"
	"database/sql"
	"encoding/pem"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/snowflakedb/gosnowflake"
	"github.com/stretchr/testify/require"
)

type StreamProps struct {
	Name    string
	Comment string
	Mode    string
}

func openSnowflake(t *testing.T) *sql.DB {
	t.Helper()

	orgName := mustEnv(t, "SNOWFLAKE_ORGANIZATION_NAME")
	accountName := mustEnv(t, "SNOWFLAKE_ACCOUNT_NAME")
	user := mustEnv(t, "SNOWFLAKE_USER")
	privateKeyPEM := mustEnv(t, "SNOWFLAKE_PRIVATE_KEY")
	role := os.Getenv("SNOWFLAKE_ROLE")

	// Parse the private key
	block, _ := pem.Decode([]byte(privateKeyPEM))
	require.NotNil(t, block, "Failed to decode PEM block from private key")

	var privateKey *rsa.PrivateKey
	var err error

	// Try PKCS8 first, then PKCS1
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		privateKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		require.NoError(t, err, "Failed to parse private key")
	} else {
		var ok bool
		privateKey, ok = key.(*rsa.PrivateKey)
		require.True(t, ok, "Private key is not RSA")
	}

	// Build account identifier: orgname-accountname
	account := fmt.Sprintf("%s-%s", orgName, accountName)

	config := gosnowflake.Config{
		Account:       account,
		User:          user,
		Authenticator: gosnowflake.AuthTypeJwt,
		PrivateKey:    privateKey,
	}

	if role != "" {
		config.Role = role
	}

	dsn, err := gosnowflake.DSN(&config)
	require.NoError(t, err, "Failed to build DSN")

	db, err := sql.Open("snowflake", dsn)
	require.NoError(t, err)
	require.NoError(t, db.Ping())
	return db
}

func setupTestTable(t *testing.T, db *sql.DB, database, schema, tableName string) {
	t.Helper()

	// Create database if not exists
	_, err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", database))
	require.NoError(t, err, "Failed to create database")

	// Create schema if not exists
	_, err = db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s.%s", database, schema))
	require.NoError(t, err, "Failed to create schema")

	// Create table
	_, err = db.Exec(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s.%s (id INT, name VARCHAR(100))", database, schema, tableName))
	require.NoError(t, err, "Failed to create table")
}

func cleanupTestTable(t *testing.T, db *sql.DB, database, schema, tableName string) {
	t.Helper()

	_, _ = db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s.%s.%s", database, schema, tableName))
}

func streamExists(t *testing.T, db *sql.DB, database, schema, streamName string) bool {
	t.Helper()

	q := fmt.Sprintf("SHOW STREAMS LIKE '%s' IN SCHEMA %s.%s;", escapeLike(streamName), database, schema)
	rows, err := db.Query(q)
	require.NoError(t, err)
	defer func() { _ = rows.Close() }()

	return rows.Next()
}

func fetchStreamProps(t *testing.T, db *sql.DB, database, schema, streamName string) StreamProps {
	t.Helper()

	q := fmt.Sprintf("SHOW STREAMS LIKE '%s' IN SCHEMA %s.%s;", escapeLike(streamName), database, schema)
	rows, err := db.Query(q)
	require.NoError(t, err)
	defer func() { _ = rows.Close() }()

	cols, err := rows.Columns()
	require.NoError(t, err)

	// Find column indices for name, comment, mode
	nameIdx, commentIdx, modeIdx := -1, -1, -1
	for i, col := range cols {
		switch col {
		case "name":
			nameIdx = i
		case "comment":
			commentIdx = i
		case "mode":
			modeIdx = i
		}
	}
	require.NotEqual(t, -1, nameIdx, "name column not found in SHOW STREAMS output")
	require.NotEqual(t, -1, commentIdx, "comment column not found in SHOW STREAMS output")

	require.True(t, rows.Next(), "No stream found matching %s", streamName)

	// Create slice to hold all column values
	values := make([]interface{}, len(cols))
	valuePtrs := make([]interface{}, len(cols))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	err = rows.Scan(valuePtrs...)
	require.NoError(t, err)

	// Extract the values we need
	getName := func(v interface{}) string {
		if v == nil {
			return ""
		}
		if s, ok := v.(string); ok {
			return s
		}
		if b, ok := v.([]byte); ok {
			return string(b)
		}
		return fmt.Sprintf("%v", v)
	}

	props := StreamProps{
		Name:    getName(values[nameIdx]),
		Comment: getName(values[commentIdx]),
	}

	if modeIdx != -1 {
		props.Mode = getName(values[modeIdx])
	}

	return props
}

func mustEnv(t *testing.T, key string) string {
	t.Helper()
	v := strings.TrimSpace(os.Getenv(key))
	require.NotEmpty(t, v, "Missing required environment variable %s", key)
	return v
}

func escapeLike(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}
