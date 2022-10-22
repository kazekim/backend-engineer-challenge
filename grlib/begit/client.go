package begit

import "github.com/kazekim/backend-engineer-challenge/grlib/befiles"

type Client interface {
	Config() *Config
	FileClient() befiles.Client

	InitGitWithSettings(name, url string) Git
}

type defaultClient struct {
	cfg *Config
	fc  befiles.Client
}

//NewClient new git client
func NewClient(cfg *Config, fc befiles.Client) Client {
	return &defaultClient{
		cfg: cfg,
		fc:  fc,
	}
}

//Config return git client config
func (c *defaultClient) Config() *Config {
	return c.cfg
}

//FileClient return file client config
func (c *defaultClient) FileClient() befiles.Client {
	return c.fc
}
