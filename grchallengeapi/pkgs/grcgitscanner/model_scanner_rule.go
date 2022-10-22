package grcgitscanner

import grcenums "github.com/kazekim/backend-engineer-challenge/grchallengeapi/enums"

type ScannerRule struct {
	Type        ScannerType       `json:"type"`
	RuleId      string            `json:"rule_id"`
	Word        string            `json:"word"`
	Description string            `json:"description"`
	Severity    grcenums.Severity `json:"severity"`
}
