package grgitrepositorydb

import (
	besqlxmocks "github.com/kazekim/backend-engineer-challenge/grlib/besqlx/mocks"
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// TestDefaultClient_GetGitRepositoryScanResultById test case for defaultClient GetGitRepositoryScanResultById
func TestDefaultClient_GetGitRepositoryScanResultById(t *testing.T) {

	rId := "xxx"
	var resp grgitrepositorydbdaos.GitRepositoryScanResult
	resp.Id = rId

	t.Run("get git repository scan result by id success", func(t *testing.T) {

		mockSqlx := new(besqlxmocks.Client)
		mockSqlx.On("Get", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(
			func(dest interface{}, query string, args ...interface{}) grerrors.Error {
				return nil
			},
		)

		c := NewClient(mockSqlx)
		m, vErr := c.GetGitRepositoryScanResultById(rId)

		assert.Nil(t, vErr)
		assert.NotNil(t, m)
	})

	t.Run("get git repository scan result by id data not found fail", func(t *testing.T) {

		mockSqlx := new(besqlxmocks.Client)
		mockSqlx.On("Get", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(
			func(dest interface{}, query string, args ...interface{}) grerrors.Error {
				return grerrors.ErrDataNotFound
			},
		)

		c := NewClient(mockSqlx)
		m, vErr := c.GetGitRepositoryScanResultById(rId)

		assert.NotNil(t, vErr)
		assert.Nil(t, m)
		assert.Equal(t, grerrors.ErrCodeDataNotFound, vErr.Code())
	})

	t.Run("get git repository scan result by id other fail", func(t *testing.T) {

		mockSqlx := new(besqlxmocks.Client)
		mockSqlx.On("Get", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(
			func(dest interface{}, query string, args ...interface{}) grerrors.Error {
				return grerrors.NewDatabaseErrorWithMessage("some error")
			},
		)

		c := NewClient(mockSqlx)
		m, vErr := c.GetGitRepositoryScanResultById(rId)

		assert.NotNil(t, vErr)
		assert.Nil(t, m)
		assert.Equal(t, grerrors.ErrCodeDatabaseError, vErr.Code())
	})
}
