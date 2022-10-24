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

//TestDefaultUseCase_CreateGitRepository test function for use case CreateGitRepository func
func TestDefaultUseCase_CreateGitRepository(t *testing.T) {

	data := challengemodels.CreateGitRepositoryData{
		Name: "test repo",
		Url: "https://github.com/kazekim/test.git",
	}

	t.Run("create git repository success", func(t *testing.T) {

		mockRepo := new(challengerepositorymocks.Repository)
		mockRepo.On("CreateGitRepository", mock.AnythingOfType("challengemodels.CreateGitRepositoryData")).Return(
			func(data challengemodels.CreateGitRepositoryData) *grmodels.GitRepository {
				return &grmodels.MockGitRepository
			},
			func(data challengemodels.CreateGitRepositoryData) grerrors.Error {
				return nil
			},
		)

		mockKafkaClient := new(grscankafkamocks.Client)

		u := NewUseCase(&grcenv.EnvMock, mockRepo, mockKafkaClient)
		resp, vErr := u.CreateGitRepository(data)

		assert.Nil(t, vErr)
		assert.NotNil(t, resp)
		assert.Equal(t, grmodels.MockGitRepository.Id, resp.Id)

	})

	t.Run("create git repository fail", func(t *testing.T) {

		mockRepo := new(challengerepositorymocks.Repository)
		mockRepo.On("CreateGitRepository", mock.AnythingOfType("challengemodels.CreateGitRepositoryData")).Return(
			func(data challengemodels.CreateGitRepositoryData) *grmodels.GitRepository {
				return nil
			},
			func(data challengemodels.CreateGitRepositoryData) grerrors.Error {
				return grerrors.ErrBadInput
			},
		)

		mockKafkaClient := new(grscankafkamocks.Client)

		u := NewUseCase(&grcenv.EnvMock, mockRepo, mockKafkaClient)
		resp, vErr := u.CreateGitRepository(data)

		assert.Nil(t, resp)
		assert.NotNil(t, vErr)

	})

}
