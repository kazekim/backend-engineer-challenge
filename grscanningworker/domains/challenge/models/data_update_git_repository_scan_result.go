package challengemodels

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grenums"
	"github.com/kazekim/backend-engineer-challenge/grlib/grgitscanner"
	"time"
)

type UpdateGitRepositoryScanResultData struct {
	Status       *grenums.ScanStatus `form:"status"`
	Findings     *grgitscanner.GitScanningResult       `form:"findings"`
	ScanningAt   *time.Time       `form:"scanning_at"`
	FinishedAt   *time.Time      `form:"finished_at"`
}