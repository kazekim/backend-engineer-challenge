package grgitscanner

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/begit"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

type ScanFileResult struct {
	Findings *[]Finding
	Error    grerrors.Error
}

// ScanGitFromGitSettings scan git with scanners from the GitSettings
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

	chanMaxSize := c.cfg.ChannelAmountPerRequest
	fq := *files

	fChan := make(chan ScanFileResult)
	var sr GitScanningResult
	for i := 0 ; i < chanMaxSize ; i++ {
		if len(fq) > 0 {
			f := fq[0]
			if len(fq) > 1 {
				fq = fq[1:]
			}else{
				fq = []string{}
			}
			go c.scanFileByLine(g, *scanners, f, fChan)
		}
	}

	count := 0
	for count < len(*files) {
		result := <-fChan
		if result.Error != nil {
			return nil, vErr
		} else {
			sr.Findings = append(sr.Findings, *result.Findings...)
		}
		count++

		if len(fq) > 0 {
			f := fq[0]
			if len(fq) > 1 {
				fq = fq[1:]
			}else{
				fq = []string{}
			}
			go c.scanFileByLine(g, *scanners, f, fChan)
		}
	}

	return &sr, nil
}

//scanFileByLine go routine func to all lines for each file with scanner
func (c *defaultClient) scanFileByLine(g begit.Git, scanners []Scanner, fileFullPath string, fChan chan ScanFileResult) {

	var result ScanFileResult
	fs := []Finding{}
	vErr := g.ScanFileByLine(fileFullPath, func(data string, line int64) grerrors.Error {

		for _, scanner := range scanners {
			input := NewScannerInput(data, fileFullPath, line)
			findings, vErr := scanner.DoScanForInput(input)
			if vErr != nil {
				result.Error = vErr
				fChan <- result
				return vErr
			}
			fs = append(fs, *findings...)
		}

		return nil
	})
	if vErr != nil {
		fChan <- result
		return
	}

	result.Findings = &fs
	fChan <- result
}
