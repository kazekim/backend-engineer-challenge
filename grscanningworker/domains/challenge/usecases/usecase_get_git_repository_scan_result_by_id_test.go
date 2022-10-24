package challengeusecases

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grgitscanner"
	grgitscannermocks "github.com/kazekim/backend-engineer-challenge/grlib/grgitscanner/mocks"
	"github.com/kazekim/backend-engineer-challenge/grlib/grmodels"
	challengemodels "github.com/kazekim/backend-engineer-challenge/grscanningworker/domains/challenge/models"
	challengerepositorymocks "github.com/kazekim/backend-engineer-challenge/grscanningworker/domains/challenge/repositories/mocks"
	grwenv "github.com/kazekim/backend-engineer-challenge/grscanningworker/env"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// TestDefaultUseCase_DoGitRepositoryScanningForResultId test function for use case DoGitRepositoryScanningForResultId func
func TestDefaultUseCase_DoGitRepositoryScanningForResultId(t *testing.T) {

	mockRepoId := "xxxxx"

	t.Run("do git repository scanning for result id success", func(t *testing.T) {

		mockRepo := new(challengerepositorymocks.Repository)
		mockRepo.On("GetGitRepositoryScanResultById", mock.AnythingOfType("string")).Return(
			func(id string) *grmodels.GitRepositoryScanResultWithDetail {
				return &grmodels.MockGitRepositoryScanResultWithDetail
			},
			func(id string) grerrors.Error {
				return nil
			},
		)
		mockRepo.On("UpdateGitRepositoryScanResultById", mock.AnythingOfType("string"), mock.AnythingOfType("challengemodels.UpdateGitRepositoryScanResultData")).Return(
			func(id string, data challengemodels.UpdateGitRepositoryScanResultData) grerrors.Error {
				return nil
			},
		)

		mockScannerClient := new(grgitscannermocks.Client)
		mockScannerClient.On("ScanGitFromGitSettings", mock.AnythingOfType("*grgitscanner.GitSettings")).Return(
			func(gs *grgitscanner.GitSettings) *grgitscanner.GitScanningResult {
				return &grgitscanner.MockGitScanningResult
			},

			func(gs *grgitscanner.GitSettings) grerrors.Error {
				return nil
			},
		)

		u := NewUseCase(&grwenv.EnvMock, mockRepo, mockScannerClient)
		vErr := u.DoGitRepositoryScanningForResultId(mockRepoId)

		assert.Nil(t, vErr)

	})

	t.Run("do git repository scanning for result id get scan result fail", func(t *testing.T) {

		mockRepo := new(challengerepositorymocks.Repository)
		mockRepo.On("GetGitRepositoryScanResultById", mock.AnythingOfType("string")).Return(
			func(id string) *grmodels.GitRepositoryScanResultWithDetail {
				return nil
			},
			func(id string) grerrors.Error {
				return grerrors.ErrInternalServerError
			},
		)
		mockRepo.On("UpdateGitRepositoryScanResultById", mock.AnythingOfType("string"), mock.AnythingOfType("challengemodels.UpdateGitRepositoryScanResultData")).Return(
			func(id string, data challengemodels.UpdateGitRepositoryScanResultData) grerrors.Error {
				return nil
			},
		)

		mockScannerClient := new(grgitscannermocks.Client)
		mockScannerClient.On("ScanGitFromGitSettings", mock.AnythingOfType("*grgitscanner.GitSettings")).Return(
			func(gs *grgitscanner.GitSettings) *grgitscanner.GitScanningResult {
				return &grgitscanner.MockGitScanningResult
			},

			func(gs *grgitscanner.GitSettings) grerrors.Error {
				return nil
			},
		)

		u := NewUseCase(&grwenv.EnvMock, mockRepo, mockScannerClient)
		vErr := u.DoGitRepositoryScanningForResultId(mockRepoId)

		assert.NotNil(t, vErr)
		assert.Equal(t, grerrors.ErrInternalServerError, vErr)

	})

	t.Run("do git repository scanning for result id update scan result fail", func(t *testing.T) {

		mockRepo := new(challengerepositorymocks.Repository)
		mockRepo.On("GetGitRepositoryScanResultById", mock.AnythingOfType("string")).Return(
			func(id string) *grmodels.GitRepositoryScanResultWithDetail {
				return &grmodels.MockGitRepositoryScanResultWithDetail
			},
			func(id string) grerrors.Error {
				return nil
			},
		)
		mockRepo.On("UpdateGitRepositoryScanResultById", mock.AnythingOfType("string"), mock.AnythingOfType("challengemodels.UpdateGitRepositoryScanResultData")).Return(
			func(id string, data challengemodels.UpdateGitRepositoryScanResultData) grerrors.Error {
				return grerrors.ErrInternalServerError
			},
		)

		mockScannerClient := new(grgitscannermocks.Client)
		mockScannerClient.On("ScanGitFromGitSettings", mock.AnythingOfType("*grgitscanner.GitSettings")).Return(
			func(gs *grgitscanner.GitSettings) *grgitscanner.GitScanningResult {
				return &grgitscanner.MockGitScanningResult
			},

			func(gs *grgitscanner.GitSettings) grerrors.Error {
				return nil
			},
		)

		u := NewUseCase(&grwenv.EnvMock, mockRepo, mockScannerClient)
		vErr := u.DoGitRepositoryScanningForResultId(mockRepoId)

		assert.NotNil(t, vErr)
		assert.Equal(t, grerrors.ErrInternalServerError, vErr)

	})

	t.Run("do git repository scanning for result id scan git from settings fail", func(t *testing.T) {

		mockRepo := new(challengerepositorymocks.Repository)
		mockRepo.On("GetGitRepositoryScanResultById", mock.AnythingOfType("string")).Return(
			func(id string) *grmodels.GitRepositoryScanResultWithDetail {
				return &grmodels.MockGitRepositoryScanResultWithDetail
			},
			func(id string) grerrors.Error {
				return nil
			},
		)
		mockRepo.On("UpdateGitRepositoryScanResultById", mock.AnythingOfType("string"), mock.AnythingOfType("challengemodels.UpdateGitRepositoryScanResultData")).Return(
			func(id string, data challengemodels.UpdateGitRepositoryScanResultData) grerrors.Error {
				return nil
			},
		)

		mockScannerClient := new(grgitscannermocks.Client)
		mockScannerClient.On("ScanGitFromGitSettings", mock.AnythingOfType("*grgitscanner.GitSettings")).Return(
			func(gs *grgitscanner.GitSettings) *grgitscanner.GitScanningResult {
				return nil
			},

			func(gs *grgitscanner.GitSettings) grerrors.Error {
				return grerrors.ErrInternalServerError
			},
		)

		u := NewUseCase(&grwenv.EnvMock, mockRepo, mockScannerClient)
		vErr := u.DoGitRepositoryScanningForResultId(mockRepoId)

		assert.Nil(t, vErr)

	})
}
