package grgitrepositorydb

import (
	"database/sql"
	besqlxmocks "github.com/kazekim/backend-engineer-challenge/grlib/besqlx/mocks"
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grenums"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

//TestDefaultClient_CreateGitRepositoryScanResult test case for defaultClient CreateGitRepositoryScanResult
func TestDefaultClient_CreateGitRepositoryScanResult(t *testing.T) {

	params := grgitrepositorydbdaos.GitRepositoryScanResult{
		RepositoryId: "xxx",
		Status: grenums.ScanStatusQueued,
		QueuedAt: sql.NullTime{
			Time: time.Now(),
			Valid: true,
		},
	}

	t.Run("create git repository scan result success", func(t *testing.T) {

		mockSqlx := new(besqlxmocks.Client)
		mockSqlx.On("NamedExec", mock.AnythingOfType("string"), mock.AnythingOfType("string"),
			mock.Anything).Return(
			func(paramQuery, valueQuery string, arg interface{}) sql.Result {
				return nil
			},
			func(paramQuery, valueQuery string, arg interface{}) grerrors.Error {
				return nil
			},
		)

		c := NewClient(mockSqlx)
		m, vErr := c.CreateGitRepositoryScanResult(&params)

		assert.Nil(t, vErr)
		assert.NotNil(t, m)
		assert.Equal(t, params.RepositoryId, m.RepositoryId)
	})

	t.Run("create git repository scan result fail", func(t *testing.T) {

		mockSqlx := new(besqlxmocks.Client)
		mockSqlx.On("NamedExec", mock.AnythingOfType("string"), mock.AnythingOfType("string"),
			mock.Anything).Return(
			func(paramQuery, valueQuery string, arg interface{}) sql.Result {
				return nil
			},
			func(paramQuery, valueQuery string, arg interface{}) grerrors.Error {
				return grerrors.ErrBadInput
			},
		)

		c := NewClient(mockSqlx)
		m, vErr := c.CreateGitRepositoryScanResult(&params)

		assert.Nil(t, m)
		assert.NotNil(t, vErr)
	})
}
