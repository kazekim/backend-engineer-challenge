package challengeusecases

import (
	challengerepositorymocks "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/repositories/mocks"
	grcenv "github.com/kazekim/backend-engineer-challenge/grchallengeapi/env"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	grscankafkamocks "github.com/kazekim/backend-engineer-challenge/grlib/grscankafka/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

//TestDefaultUseCase_DeleteGitRepositoryById test function for use case DeleteGitRepositoryById func
func TestDefaultUseCase_DeleteGitRepositoryById(t *testing.T) {

	mockId := "xxxxx"

	t.Run("delete git repository by id success", func(t *testing.T) {

		mockRepo := new(challengerepositorymocks.Repository)
		mockRepo.On("DeleteGitRepositoryById", mock.AnythingOfType("string")).Return(
			func(id string) grerrors.Error {
				return nil
			},
		)

		mockKafkaClient := new(grscankafkamocks.Client)

		u := NewUseCase(&grcenv.EnvMock, mockRepo, mockKafkaClient)
		vErr := u.DeleteGitRepositoryById(mockId)

		assert.Nil(t, vErr)

	})

	t.Run("delete git repository by id fail", func(t *testing.T) {

		mockRepo := new(challengerepositorymocks.Repository)
		mockRepo.On("DeleteGitRepositoryById", mock.AnythingOfType("string")).Return(
			func(id string) grerrors.Error {
				return grerrors.ErrBadInput
			},
		)

		mockKafkaClient := new(grscankafkamocks.Client)

		u := NewUseCase(&grcenv.EnvMock, mockRepo, mockKafkaClient)
		vErr := u.DeleteGitRepositoryById(mockId)

		assert.NotNil(t, vErr)
		assert.Equal(t, grerrors.ErrBadInput, vErr)

	})

}
