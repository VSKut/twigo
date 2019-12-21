package repository

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Constructor(t *testing.T) {
	t.Run("NewPostgreSQLRepository", func(t *testing.T) {
		repo := NewPostgreSQLRepository(&sql.DB{})

		assert.IsType(t, &Repository{}, repo, "should be repository")
	})

}
