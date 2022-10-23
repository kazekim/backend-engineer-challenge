package grgitscanner

import (
	"fmt"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

//DoScanForInput start scan for input with rule that specify
func (s *defaultScanner) DoScanForInput(input ScannerInput) (*[]GitScanningResult, grerrors.Error) {

	var gsrs []GitScanningResult
	switch s.rule.Type {
	case ScannerTypeBeginWord:
		isFound, result, vErr := s.detectBeginWithWord(input)
		if vErr != nil {
			return nil, vErr
		}
		if isFound {
			gsrs = append(gsrs, *result)
		}
	default:
		msg := fmt.Sprintf("%v: invalid scanner type '%v'", packageName, s.rule.Type)
		return nil, grerrors.NewDefaultErrorWithMessage(msg)
	}

	return &gsrs, nil
}
