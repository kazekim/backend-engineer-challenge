package grgitrepositorydb

import (
	besqlxmocks "github.com/kazekim/backend-engineer-challenge/grlib/besqlx/mocks"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

//TestDefaultClient_DeleteGitRepositoryById test case for defaultClient DeleteGitRepositoryById
func TestDefaultClient_DeleteGitRepositoryById(t *testing.T) {

	rId := "xxx"

	t.Run("delete git repository by id success", func(t *testing.T) {

		mockSqlx := new(besqlxmocks.Client)
		mockSqlx.On("DeleteExec", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(
			func(arg interface{}, whereQuery string, values ...interface{}) grerrors.Error {
				return nil
			},
		)

		c := NewClient(mockSqlx)
		vErr := c.DeleteGitRepositoryById(rId)

		assert.Nil(t, vErr)
	})

	t.Run("delete git repository by id fail", func(t *testing.T) {

		mockSqlx := new(besqlxmocks.Client)
		mockSqlx.On("DeleteExec", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(
			func(arg interface{}, whereQuery string, values ...interface{}) grerrors.Error {
				return grerrors.ErrInternalServerError
			},
		)

		c := NewClient(mockSqlx)
		vErr := c.DeleteGitRepositoryById(rId)

		assert.NotNil(t, vErr)
	})
}
