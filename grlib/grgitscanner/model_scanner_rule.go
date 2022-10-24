package grgitscanner

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grenums"
)

type ScannerRule struct {
	Type        ScannerType      `json:"type"`
	RuleId      string           `json:"rule_id"`
	Word        string           `json:"word"`
	Description string           `json:"description"`
	Severity    grenums.Severity `json:"severity"`
}
