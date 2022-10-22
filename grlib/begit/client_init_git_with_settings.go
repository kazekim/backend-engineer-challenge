package begit

func (c *defaultClient) InitGitWithSettings(name, url string) *Git {
	return &Git{
		c:    c,
		name: name,
		url:  url,
	}
}
