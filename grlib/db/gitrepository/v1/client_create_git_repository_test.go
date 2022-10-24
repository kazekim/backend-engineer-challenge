package grgitrepositorydb

import (
	"database/sql"
	besqlxmocks "github.com/kazekim/backend-engineer-challenge/grlib/besqlx/mocks"
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDefaultClient_CreateGitRepository(t *testing.T) {

	params := grgitrepositorydbdaos.GitRepository{
		Name:     "xxx",
		Url:      "url",
		IsActive: true,
	}

	t.Run("create git repository success", func(t *testing.T) {

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
		m, vErr := c.CreateGitRepository(&params)

		assert.Nil(t, vErr)
		assert.NotNil(t, m)
		assert.Equal(t, params.Name, m.Name)
	})

	t.Run("create git repository fail", func(t *testing.T) {

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
		m, vErr := c.CreateGitRepository(&params)

		assert.Nil(t, m)
		assert.NotNil(t, vErr)
	})
}
