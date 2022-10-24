package grgitscanner

import "github.com/kazekim/backend-engineer-challenge/grlib/grerrors"

//LoadScanners load default scanner
//can be modified to load scanner from database in the future
func (c *defaultClient) LoadScanners() (*[]Scanner, grerrors.Error) {

	var ss []Scanner
	for _, r := range rules {
		scanner := NewScanner(r)
		ss = append(ss, scanner)
	}
	return &ss, nil
}
