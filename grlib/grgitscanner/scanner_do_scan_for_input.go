package grgitscanner

import (
	"fmt"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

// DoScanForInput start scan for input with rule that specify
func (s *defaultScanner) DoScanForInput(input ScannerInput) (*[]Finding, grerrors.Error) {

	var fs []Finding
	switch s.rule.Type {
	case ScannerTypeBeginWord:
		isFound, finding, vErr := s.detectBeginWithWord(input)
		if vErr != nil {
			return nil, vErr
		}
		if isFound {
			fs = append(fs, *finding)
		}
	default:
		msg := fmt.Sprintf("%v: invalid scanner type '%v'", packageName, s.rule.Type)
		return nil, grerrors.NewDefaultErrorWithMessage(msg)
	}

	return &fs, nil
}
