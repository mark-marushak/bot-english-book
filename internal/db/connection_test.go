package db

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDB(t *testing.T) {
	t.Run("Test Connection to DB", func(t *testing.T) {
		t.Parallel()

		db := connect()
		assert.IsTypef(t, &sql.DB{}, db, "variable db is not type of *gorm.DB")
	})
}
