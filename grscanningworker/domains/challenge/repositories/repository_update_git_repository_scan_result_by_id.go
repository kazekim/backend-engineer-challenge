package challengerepositories

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	challengemodels "github.com/kazekim/backend-engineer-challenge/grscanningworker/domains/challenge/models"
)

// UpdateGitRepositoryScanResultById update git repository scan result by result id and update datas with db client
func (r *defaultRepository) UpdateGitRepositoryScanResultById(id string, data challengemodels.UpdateGitRepositoryScanResultData) grerrors.Error {

	uDb := r.grc.UpdateGitRepositoryScanResultById(id)
	if data.Status != nil {
		uDb.Status(*data.Status)
	}
	if data.Findings != nil {
		uDb.Findings(*data.Findings)
	}
	if data.ScanningAt != nil {
		uDb.ScanningAt(*data.ScanningAt)
	}
	if data.FinishedAt != nil {
		uDb.FinishedAt(*data.FinishedAt)
	}
	vErr := uDb.Commit()
	if vErr != nil {
		return vErr
	}

	return nil
}
