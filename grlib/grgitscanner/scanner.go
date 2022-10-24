package grgitscanner

import "github.com/kazekim/backend-engineer-challenge/grlib/grerrors"

type Scanner interface {
	DoScanForInput(input ScannerInput) (*[]Finding, grerrors.Error)
}

type defaultScanner struct {
	rule ScannerRule
}

// NewScanner new git scanner with rules
func NewScanner(rule ScannerRule) Scanner {

	s := defaultScanner{
		rule: rule,
	}
	return &s
}
