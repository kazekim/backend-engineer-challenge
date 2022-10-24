package grgitrepositorydb

import (
	besqlxmocks "github.com/kazekim/backend-engineer-challenge/grlib/besqlx/mocks"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

//TestDefaultClient_DeleteGitRepositoryScanResultById test case for defaultClient DeleteGitRepositoryScanResultById
func TestDefaultClient_DeleteGitRepositoryScanResultById(t *testing.T) {

	rId := "xxx"

	t.Run("delete git repository scan result by id success", func(t *testing.T) {

		mockSqlx := new(besqlxmocks.Client)
		mockSqlx.On("DeleteExec", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(
			func(arg interface{}, whereQuery string, values ...interface{}) grerrors.Error {
				return nil
			},
		)

		c := NewClient(mockSqlx)
		vErr := c.DeleteGitRepositoryScanResultById(rId)

		assert.Nil(t, vErr)
	})

	t.Run("delete git repository scan result by id fail", func(t *testing.T) {

		mockSqlx := new(besqlxmocks.Client)
		mockSqlx.On("DeleteExec", mock.Anything, mock.AnythingOfType("string"), mock.Anything).Return(
			func(arg interface{}, whereQuery string, values ...interface{}) grerrors.Error {
				return grerrors.ErrInternalServerError
			},
		)

		c := NewClient(mockSqlx)
		vErr := c.DeleteGitRepositoryScanResultById(rId)

		assert.NotNil(t, vErr)
	})
}
