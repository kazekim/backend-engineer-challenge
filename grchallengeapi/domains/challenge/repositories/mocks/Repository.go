// Code generated by mockery v2.14.0. DO NOT EDIT.

package challengerepositorymocks

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"

	grerrors "github.com/kazekim/backend-engineer-challenge/grlib/grerrors"

	grmodels "github.com/kazekim/backend-engineer-challenge/grlib/grmodels"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateGitRepository provides a mock function with given fields: data
func (_m *Repository) CreateGitRepository(data challengemodels.CreateGitRepositoryData) (*grmodels.GitRepository, grerrors.Error) {
	ret := _m.Called(data)

	var r0 *grmodels.GitRepository
	if rf, ok := ret.Get(0).(func(challengemodels.CreateGitRepositoryData) *grmodels.GitRepository); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grmodels.GitRepository)
		}
	}

	var r1 grerrors.Error
	if rf, ok := ret.Get(1).(func(challengemodels.CreateGitRepositoryData) grerrors.Error); ok {
		r1 = rf(data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(grerrors.Error)
		}
	}

	return r0, r1
}

// DeleteGitRepositoryById provides a mock function with given fields: id
func (_m *Repository) DeleteGitRepositoryById(id string) grerrors.Error {
	ret := _m.Called(id)

	var r0 grerrors.Error
	if rf, ok := ret.Get(0).(func(string) grerrors.Error); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(grerrors.Error)
		}
	}

	return r0
}

// DeleteGitRepositoryScanResultById provides a mock function with given fields: id
func (_m *Repository) DeleteGitRepositoryScanResultById(id string) grerrors.Error {
	ret := _m.Called(id)

	var r0 grerrors.Error
	if rf, ok := ret.Get(0).(func(string) grerrors.Error); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(grerrors.Error)
		}
	}

	return r0
}

// GetGitRepositoryById provides a mock function with given fields: id
func (_m *Repository) GetGitRepositoryById(id string) (*grmodels.GitRepository, grerrors.Error) {
	ret := _m.Called(id)

	var r0 *grmodels.GitRepository
	if rf, ok := ret.Get(0).(func(string) *grmodels.GitRepository); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grmodels.GitRepository)
		}
	}

	var r1 grerrors.Error
	if rf, ok := ret.Get(1).(func(string) grerrors.Error); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(grerrors.Error)
		}
	}

	return r0, r1
}

// GetGitRepositoryScanResultById provides a mock function with given fields: id
func (_m *Repository) GetGitRepositoryScanResultById(id string) (*grmodels.GitRepositoryScanResultWithDetail, grerrors.Error) {
	ret := _m.Called(id)

	var r0 *grmodels.GitRepositoryScanResultWithDetail
	if rf, ok := ret.Get(0).(func(string) *grmodels.GitRepositoryScanResultWithDetail); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grmodels.GitRepositoryScanResultWithDetail)
		}
	}

	var r1 grerrors.Error
	if rf, ok := ret.Get(1).(func(string) grerrors.Error); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(grerrors.Error)
		}
	}

	return r0, r1
}

// ListGitRepositories provides a mock function with given fields: filter, values
func (_m *Repository) ListGitRepositories(filter challengemodels.GitRepositoryFilterData, values ...int64) (*[]grmodels.GitRepository, int64, grerrors.Error) {
	_va := make([]interface{}, len(values))
	for _i := range values {
		_va[_i] = values[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *[]grmodels.GitRepository
	if rf, ok := ret.Get(0).(func(challengemodels.GitRepositoryFilterData, ...int64) *[]grmodels.GitRepository); ok {
		r0 = rf(filter, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]grmodels.GitRepository)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(challengemodels.GitRepositoryFilterData, ...int64) int64); ok {
		r1 = rf(filter, values...)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 grerrors.Error
	if rf, ok := ret.Get(2).(func(challengemodels.GitRepositoryFilterData, ...int64) grerrors.Error); ok {
		r2 = rf(filter, values...)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(grerrors.Error)
		}
	}

	return r0, r1, r2
}

// ListGitRepositoryScanResults provides a mock function with given fields: filter, values
func (_m *Repository) ListGitRepositoryScanResults(filter challengemodels.GitRepositoryScanResultFilterData, values ...int64) (*[]grmodels.GitRepositoryScanResultWithDetail, int64, grerrors.Error) {
	_va := make([]interface{}, len(values))
	for _i := range values {
		_va[_i] = values[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, filter)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *[]grmodels.GitRepositoryScanResultWithDetail
	if rf, ok := ret.Get(0).(func(challengemodels.GitRepositoryScanResultFilterData, ...int64) *[]grmodels.GitRepositoryScanResultWithDetail); ok {
		r0 = rf(filter, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]grmodels.GitRepositoryScanResultWithDetail)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(challengemodels.GitRepositoryScanResultFilterData, ...int64) int64); ok {
		r1 = rf(filter, values...)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 grerrors.Error
	if rf, ok := ret.Get(2).(func(challengemodels.GitRepositoryScanResultFilterData, ...int64) grerrors.Error); ok {
		r2 = rf(filter, values...)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(grerrors.Error)
		}
	}

	return r0, r1, r2
}

// StartGitRepositoryScanning provides a mock function with given fields: id
func (_m *Repository) StartGitRepositoryScanning(id string) (*grmodels.GitRepositoryScanResult, grerrors.Error) {
	ret := _m.Called(id)

	var r0 *grmodels.GitRepositoryScanResult
	if rf, ok := ret.Get(0).(func(string) *grmodels.GitRepositoryScanResult); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*grmodels.GitRepositoryScanResult)
		}
	}

	var r1 grerrors.Error
	if rf, ok := ret.Get(1).(func(string) grerrors.Error); ok {
		r1 = rf(id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(grerrors.Error)
		}
	}

	return r0, r1
}

// UpdateGitRepositoryById provides a mock function with given fields: id, data
func (_m *Repository) UpdateGitRepositoryById(id string, data challengemodels.UpdateGitRepositoryData) grerrors.Error {
	ret := _m.Called(id, data)

	var r0 grerrors.Error
	if rf, ok := ret.Get(0).(func(string, challengemodels.UpdateGitRepositoryData) grerrors.Error); ok {
		r0 = rf(id, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(grerrors.Error)
		}
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
