package grgitscanner

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

//ScanGitFromGitSettings scan git with scanners from the GitSettings
func (c *defaultClient) ScanGitFromGitSettings(gs *GitSettings) (*GitScanningResult, grerrors.Error) {

	g := c.gc.InitGitWithSettings(gs.Name, gs.Url)
	defer func() {
		_ = g.Close()
	}()

	vErr := g.Clone()
	if vErr != nil {
		return nil, vErr
	}

	files, vErr := g.ListFiles()
	if vErr != nil {
		return nil, vErr
	}

	scanners, vErr := c.LoadScanners()
	if vErr != nil {
		return nil, vErr
	}

	var sr GitScanningResult
	for _, f := range *files {
		vErr = g.ScanFileByLine(f, func(data string, line int64) grerrors.Error {

			for _, scanner := range *scanners {
				input := NewScannerInput(data, f, line)
				findings, vErr := scanner.DoScanForInput(input)
				if vErr != nil {
					return vErr
				}
				sr.Findings = append(sr.Findings, (*findings)...)
			}

			return nil
		})
		if vErr != nil {
			return nil, vErr
		}
	}

	return &sr, nil
}
