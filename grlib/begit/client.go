package begit

type Client interface {
	Config() *Config

	InitGitWithSettings(name, url string) *Git
}

type defaultClient struct {
	cfg *Config
}

func NewClient(cfg *Config) Client {
	return &defaultClient{
		cfg: cfg,
	}
}

func (c *defaultClient) Config() *Config {
	return c.cfg
}
