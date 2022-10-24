package challengeusecases

import (
	challengerepositorymocks "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/repositories/mocks"
	grcenv "github.com/kazekim/backend-engineer-challenge/grchallengeapi/env"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
	grscankafkamocks "github.com/kazekim/backend-engineer-challenge/grlib/grscankafka/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

//TestDefaultUseCase_GetGitRepositoryScanResultById test function for use case GetGitRepositoryScanResultById func
func TestDefaultUseCase_GetGitRepositoryScanResultById(t *testing.T) {

	mockId := "xxxxx"

	t.Run("get git repository scan result by id success", func(t *testing.T) {

		mockRepo := new(challengerepositorymocks.Repository)
		mockRepo.On("GetGitRepositoryScanResultById", mock.AnythingOfType("string")).Return(
			func(id string) *grmodels.GitRepositoryScanResultWithDetail {
				return &grmodels.MockGitRepositoryScanResultWithDetail
			},
			func(id string) grerrors.Error {
				return nil
			},
		)

		mockKafkaClient := new(grscankafkamocks.Client)

		u := NewUseCase(&grcenv.EnvMock, mockRepo, mockKafkaClient)
		m, vErr := u.GetGitRepositoryScanResultById(mockId)

		assert.Nil(t, vErr)
		assert.NotNil(t, m)
		assert.Equal(t, grmodels.MockGitRepositoryScanResult.Id, m.Id)

	})

	t.Run("get git repository scan result by id fail", func(t *testing.T) {

		mockRepo := new(challengerepositorymocks.Repository)
		mockRepo.On("GetGitRepositoryScanResultById", mock.AnythingOfType("string")).Return(
			func(id string) *grmodels.GitRepositoryScanResultWithDetail {
				return nil
			},
			func(id string) grerrors.Error {
				return grerrors.ErrBadInput
			},
		)

		mockKafkaClient := new(grscankafkamocks.Client)

		u := NewUseCase(&grcenv.EnvMock, mockRepo, mockKafkaClient)
		m, vErr := u.GetGitRepositoryScanResultById(mockId)

		assert.Nil(t, m)
		assert.NotNil(t, vErr)
		assert.Equal(t, grerrors.ErrBadInput, vErr)

	})

}
