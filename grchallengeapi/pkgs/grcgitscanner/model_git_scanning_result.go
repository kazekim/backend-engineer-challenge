package grcgitscanner

import grcenums "github.com/kazekim/backend-engineer-challenge/grchallengeapi/enums"

type GitScanningResult struct {
	Type ScannerType `json:"type"`
	RuleId string `json:"rule_id"`
	Location Location `json:"location"`
	Metadata Metadata `json:"metadata"`
}

type Location struct {
	Path string `json:"path"`
	Positions Positions `json:"positions"`
}

type Positions struct {
	Begin Begin `json:"begin"`
}

type Begin struct {
	Line int64 `json:"line"`
}

type Metadata struct {
	Description string `json:"description"`
	Severity grcenums.Severity `json:"severity"`
}