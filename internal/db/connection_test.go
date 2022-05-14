package db

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestDB(t *testing.T) {
	t.Run("Test Connection to DB", func(t *testing.T) {
		t.Parallel()

		db := DB()
		assert.IsTypef(t, &gorm.DB{}, db, "variable db is not type of *gorm.DB")
	})
}
