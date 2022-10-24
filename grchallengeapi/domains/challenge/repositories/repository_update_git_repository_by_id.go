package challengerepositories

import (
	challengemodels "github.com/kazekim/backend-engineer-challenge/grchallengeapi/domains/challenge/models"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

// UpdateGitRepositoryById update git repository by repository id and update datas with db client
func (r *defaultRepository) UpdateGitRepositoryById(id string, data challengemodels.UpdateGitRepositoryData) grerrors.Error {

	uDb := r.grc.UpdateGitRepositoryById(id)
	if data.Name != nil {
		uDb.Name(*data.Name)
	}
	if data.Url != nil {
		uDb.Url(*data.Url)
	}
	if data.IsActive != nil {
		uDb.IsActive(*data.IsActive)
	}
	vErr := uDb.Commit()
	if vErr != nil {
		return vErr
	}

	return nil
}
