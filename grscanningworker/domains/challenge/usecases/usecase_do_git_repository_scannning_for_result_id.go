package challengeusecases

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grenums"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"github.com/kazekim/backend-engineer-challenge/grlib/grgitscanner"
	challengemodels "github.com/kazekim/backend-engineer-challenge/grscanningworker/domains/challenge/models"
	"time"
)

func (u *defaultUseCase) DoGitRepositoryScanningForResultId(resultId string) grerrors.Error {

	gr, vErr := u.cr.GetGitRepositoryScanResultById(resultId)
	if vErr != nil {
		return vErr
	}

	if gr.Status != grenums.ScanStatusQueued {
		return nil
	}

	settings := &grgitscanner.GitSettings{
		Name: gr.RepositoryName,
		Url: gr.RepositoryUrl,
	}

	status := grenums.ScanStatusInProgress
	t := time.Now()
	uData := challengemodels.UpdateGitRepositoryScanResultData{
		Status: &status,
		ScanningAt: &t,
	}
	vErr = u.cr.UpdateGitRepositoryScanResultById(resultId, uData)
	if vErr != nil {
		return vErr
	}

	srs, vErr := u.gsc.ScanGitFromGitSettings(settings)
	t = time.Now()
	uData.FinishedAt = &t
	if vErr != nil {
		status =  grenums.ScanStatusFailure
	}else{
		status = grenums.ScanStatusSuccess
		uData.Findings = srs
	}
	uData.Status = &status

	vErr = u.cr.UpdateGitRepositoryScanResultById(resultId, uData)
	if vErr != nil {
		return vErr
	}

	return nil
}
