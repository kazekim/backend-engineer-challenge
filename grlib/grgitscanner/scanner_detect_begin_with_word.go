package grgitscanner

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
	"strings"
)

// detectBeginWithWord scan data inline with begin with word rules and return results
func (s *defaultScanner) detectBeginWithWord(input ScannerInput) (bool, *Finding, grerrors.Error) {

	isMatched := strings.HasPrefix(input.Data, s.rule.Word)
	if isMatched {

		finding := newGitScannerResultForScannerTypeBeginWithWord(s.rule, input)
		return true, &finding, nil
	}

	return false, nil, nil
}

// newGitScannerResultForScannerTypeBeginWithWord init GitScannerResult from rule and input object
func newGitScannerResultForScannerTypeBeginWithWord(rule ScannerRule, input ScannerInput) Finding {
	result := Finding{
		Type:   rule.Type,
		RuleId: rule.RuleId,
		Location: Location{
			Path: input.FilePath,
			Positions: Positions{
				Begin: Begin{
					Line: input.Line,
				},
			},
		},
		Metadata: Metadata{
			Description: rule.Description,
			Severity:    rule.Severity,
		},
	}
	return result
}
