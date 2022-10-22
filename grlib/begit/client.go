package begit

type Client interface {
	Config() *Config

	InitGitWithSettings(name, url string) *Git
}

type defaultClient struct {
	cfg *Config
}

//NewClient new git client
func NewClient(cfg *Config) Client {
	return &defaultClient{
		cfg: cfg,
	}
}

//Config return git client config
func (c *defaultClient) Config() *Config {
	return c.cfg
}
