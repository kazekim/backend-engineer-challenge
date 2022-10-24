package challengeusecases

import (
	challengerepositorymocks "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/repositories/mocks"
	grcenv "github.com/kazekim/backend-engineer-challenge/grchallengeapi/env"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
	"github.com/kazekim/backend-engineer-challenge/grlib/grscankafka"
	grscankafkamocks "github.com/kazekim/backend-engineer-challenge/grlib/grscankafka/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// TestDefaultUseCase_StartGitRepositoryScanning test function for use case StartGitRepositoryScanning func
func TestDefaultUseCase_StartGitRepositoryScanning(t *testing.T) {

	id := "xxxx"

	t.Run("start git repository scanning success", func(t *testing.T) {

		mockRepo := new(challengerepositorymocks.Repository)
		mockRepo.On("StartGitRepositoryScanning", mock.AnythingOfType("string")).Return(
			func(id string) *grmodels.GitRepositoryScanResult {
				return &grmodels.MockGitRepositoryScanResult
			},
			func(id string) grerrors.Error {
				return nil
			},
		)

		mockKafkaClient := new(grscankafkamocks.Client)
		mockKafkaClient.On("PublishStartGitRepositoryScanningMessage", mock.AnythingOfType("string"), mock.AnythingOfType("grscankafka.PublishStartGitRepositoryScanningMessageParams")).Return(
			func(serverId string, data grscankafka.PublishStartGitRepositoryScanningMessageParams) grerrors.Error {
				return nil
			},
		)

		u := NewUseCase(&grcenv.EnvMock, mockRepo, mockKafkaClient)
		m, vErr := u.StartGitRepositoryScanning(id)

		assert.Nil(t, vErr)
		assert.NotNil(t, m)
		assert.Equal(t, grmodels.MockGitRepositoryScanResult.Id, m.Id)

	})

	t.Run("start git repository scanning create database fail", func(t *testing.T) {

		mockRepo := new(challengerepositorymocks.Repository)
		mockRepo.On("StartGitRepositoryScanning", mock.AnythingOfType("string")).Return(
			func(id string) *grmodels.GitRepositoryScanResult {
				return nil
			},
			func(id string) grerrors.Error {
				return grerrors.ErrBadInput
			},
		)

		mockKafkaClient := new(grscankafkamocks.Client)
		mockKafkaClient.On("PublishStartGitRepositoryScanningMessage", mock.AnythingOfType("string"), mock.AnythingOfType("grscankafka.PublishStartGitRepositoryScanningMessageParams")).Return(
			func(serverId string, data grscankafka.PublishStartGitRepositoryScanningMessageParams) grerrors.Error {
				return nil
			},
		)

		u := NewUseCase(&grcenv.EnvMock, mockRepo, mockKafkaClient)
		m, vErr := u.StartGitRepositoryScanning(id)

		assert.Nil(t, m)
		assert.NotNil(t, vErr)

	})

	t.Run("start git repository scanning send message to queue fail", func(t *testing.T) {

		mockRepo := new(challengerepositorymocks.Repository)
		mockRepo.On("StartGitRepositoryScanning", mock.AnythingOfType("string")).Return(
			func(id string) *grmodels.GitRepositoryScanResult {
				return &grmodels.MockGitRepositoryScanResult
			},
			func(id string) grerrors.Error {
				return nil
			},
		)

		mockKafkaClient := new(grscankafkamocks.Client)
		mockKafkaClient.On("PublishStartGitRepositoryScanningMessage", mock.AnythingOfType("string"), mock.AnythingOfType("grscankafka.PublishStartGitRepositoryScanningMessageParams")).Return(
			func(serverId string, data grscankafka.PublishStartGitRepositoryScanningMessageParams) grerrors.Error {
				return grerrors.ErrInternalServerError
			},
		)

		u := NewUseCase(&grcenv.EnvMock, mockRepo, mockKafkaClient)
		m, vErr := u.StartGitRepositoryScanning(id)

		assert.Nil(t, m)
		assert.NotNil(t, vErr)

	})

}
