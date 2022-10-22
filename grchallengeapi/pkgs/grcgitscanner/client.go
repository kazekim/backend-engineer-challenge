package grcgitscanner

import (
	"github.com/kazekim/backend-engineer-challenge/grlib/befiles"
	"github.com/kazekim/backend-engineer-challenge/grlib/begit"
	"github.com/kazekim/backend-engineer-challenge/grlib/grerrors"
)

const (
	packageName = "grcgitscanner"
)

type Client interface {
	ScanGitFromGitSettings(gs *GitSettings) (*[]GitScanningResult, grerrors.Error)
}

type defaultClient struct {
	cfg *Config
	gc begit.Client
}

//NewClient new git scanner client
func NewClient(cfg *Config, fc befiles.Client) Client {

	gc := begit.NewClient(cfg.GitConfig(), fc)

	return &defaultClient{
		cfg: cfg,
		gc: gc,
	}
}
