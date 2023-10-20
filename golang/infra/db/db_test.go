package db_test

import (
	"os"
	"testing"

	"github.com/paulojr-eco/full-cycle-desafio/infra/db"
	"github.com/stretchr/testify/require"
)

/* Should return an error if the connection url is wrong */
func TestDb_NewWrongConection(t *testing.T) {
	dbType := os.Getenv("dbType")
	t.Setenv("dbType", "wrong_database_connection")
	database, err := db.ConnectDB()

	require.NotNil(t, err)
	require.Nil(t, database)

	t.Cleanup(func() { os.Setenv("dbType", dbType) })
}

func TestDb_NewConection(t *testing.T) {
	database, err := db.ConnectDB()

	require.Nil(t, err)
	require.NotNil(t, database)
}
