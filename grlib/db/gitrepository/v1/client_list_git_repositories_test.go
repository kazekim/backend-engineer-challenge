package grgitrepositorydb

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/besqlx"
	besqlxmocks "github.com/kazekim/backend-engineer-challenge/grlib/besqlx/mocks"
	grgitrepositorydbdaos "github.com/kazekim/backend-engineer-challenge/grlib/db/gitrepository/v1/daos"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

//TestDefaultClient_ListGitRepositoriesById test case for defaultClient ListGitRepositoriesById
func TestDefaultClient_ListGitRepositoriesById(t *testing.T) {

	filter := grgitrepositorydbdaos.GitRepositoryFilter{}

	t.Run("list git repositories success", func(t *testing.T) {

		mockSqlx := new(besqlxmocks.Client)
		mockSqlx.On("SelectWithCount", mock.Anything, mock.Anything,
			mock.AnythingOfType("besqlx.SelectQueryBuilder"), mock.Anything, mock.Anything).Return(
			func(model interface{}, dest interface{}, qb besqlx.SelectQueryBuilder, values ...int64) int64 {
				return 1
			},
			func(model interface{}, dest interface{}, qb besqlx.SelectQueryBuilder, values ...int64) grerrors.Error {
				return nil
			},
		)

		c := NewClient(mockSqlx)
		ms, count, vErr := c.ListGitRepositories(filter)

		assert.Nil(t, vErr)
		assert.NotNil(t, ms)
		assert.Equal(t, int64(1), count)
	})

	t.Run("list git repositories fail", func(t *testing.T) {

		mockSqlx := new(besqlxmocks.Client)
		mockSqlx.On("SelectWithCount", mock.Anything, mock.Anything,
			mock.AnythingOfType("besqlx.SelectQueryBuilder"), mock.Anything, mock.Anything).Return(
			func(model interface{}, dest interface{}, qb besqlx.SelectQueryBuilder, values ...int64) int64 {
				return 0
			},
			func(model interface{}, dest interface{}, qb besqlx.SelectQueryBuilder, values ...int64) grerrors.Error {
				return grerrors.NewDatabaseErrorWithMessage("some error")
			},
		)

		c := NewClient(mockSqlx)
		ms, count, vErr := c.ListGitRepositories(filter)

		assert.NotNil(t, vErr)
		assert.Nil(t, ms)
		assert.Equal(t, int64(0), count)
	})


}
