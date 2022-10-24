package challengeusecases

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	challengerepositorymocks "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/repositories/mocks"
	grcenv "github.com/kazekim/backend-engineer-challenge/grchallengeapi/env"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
	grscankafkamocks "github.com/kazekim/backend-engineer-challenge/grlib/grscankafka/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// TestDefaultUseCase_ListGitRepositories test function for use case ListGitRepositories func
func TestDefaultUseCase_ListGitRepositories(t *testing.T) {

	filter := challengemodels.GitRepositoryFilterData{}

	t.Run("list git repositories success", func(t *testing.T) {

		mockRepo := new(challengerepositorymocks.Repository)
		mockRepo.On("ListGitRepositories", mock.AnythingOfType("challengemodels.GitRepositoryFilterData")).Return(
			func(filter challengemodels.GitRepositoryFilterData, values ...int64) *[]grmodels.GitRepository {
				return &[]grmodels.GitRepository{
					grmodels.MockGitRepository,
				}
			},
			func(filter challengemodels.GitRepositoryFilterData, values ...int64) int64 {
				return 1
			},
			func(filter challengemodels.GitRepositoryFilterData, values ...int64) grerrors.Error {
				return nil
			},
		)

		mockKafkaClient := new(grscankafkamocks.Client)

		u := NewUseCase(&grcenv.EnvMock, mockRepo, mockKafkaClient)
		ms, count, vErr := u.ListGitRepositories(filter)

		assert.Nil(t, vErr)
		assert.NotNil(t, ms)
		assert.Equal(t, int64(1), count)
		assert.Len(t, *ms, 1)

	})

	t.Run("list git repositories fail", func(t *testing.T) {

		mockRepo := new(challengerepositorymocks.Repository)
		mockRepo.On("ListGitRepositories", mock.AnythingOfType("challengemodels.GitRepositoryFilterData")).Return(
			func(filter challengemodels.GitRepositoryFilterData, values ...int64) *[]grmodels.GitRepository {
				return nil
			},
			func(filter challengemodels.GitRepositoryFilterData, values ...int64) int64 {
				return 0
			},
			func(filter challengemodels.GitRepositoryFilterData, values ...int64) grerrors.Error {
				return grerrors.ErrBadInput
			},
		)

		mockKafkaClient := new(grscankafkamocks.Client)

		u := NewUseCase(&grcenv.EnvMock, mockRepo, mockKafkaClient)
		ms, count, vErr := u.ListGitRepositories(filter)

		assert.Nil(t, ms)
		assert.NotNil(t, vErr)
		assert.Equal(t, int64(0), count)
		assert.Equal(t, grerrors.ErrBadInput, vErr)

	})

}
