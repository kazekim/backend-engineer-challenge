package challengeusecases

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	challengerepositorymocks "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/repositories/mocks"
	grcenv "github.com/kazekim/backend-engineer-challenge/grchallengeapi/env"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	grscankafkamocks "github.com/kazekim/backend-engineer-challenge/grlib/grscankafka/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// TestDefaultUseCase_UpdateGitRepositoryById test function for use case UpdateGitRepositoryById func
func TestDefaultUseCase_UpdateGitRepositoryById(t *testing.T) {

	mockId := "xxxxx"
	mockUpdateData := challengemodels.UpdateGitRepositoryData{}

	t.Run("update git repository by id success", func(t *testing.T) {

		mockRepo := new(challengerepositorymocks.Repository)
		mockRepo.On("UpdateGitRepositoryById", mock.AnythingOfType("string"), mock.AnythingOfType("challengemodels.UpdateGitRepositoryData")).Return(
			func(id string, data challengemodels.UpdateGitRepositoryData) grerrors.Error {
				return nil
			},
		)

		mockKafkaClient := new(grscankafkamocks.Client)

		u := NewUseCase(&grcenv.EnvMock, mockRepo, mockKafkaClient)
		vErr := u.UpdateGitRepositoryById(mockId, mockUpdateData)

		assert.Nil(t, vErr)

	})

	t.Run("update git repository by id fail", func(t *testing.T) {

		mockRepo := new(challengerepositorymocks.Repository)
		mockRepo.On("UpdateGitRepositoryById", mock.AnythingOfType("string"), mock.AnythingOfType("challengemodels.UpdateGitRepositoryData")).Return(
			func(id string, data challengemodels.UpdateGitRepositoryData) grerrors.Error {
				return grerrors.ErrBadInput
			},
		)

		mockKafkaClient := new(grscankafkamocks.Client)

		u := NewUseCase(&grcenv.EnvMock, mockRepo, mockKafkaClient)
		vErr := u.UpdateGitRepositoryById(mockId, mockUpdateData)

		assert.NotNil(t, vErr)
		assert.Equal(t, grerrors.ErrBadInput, vErr)

	})

}
