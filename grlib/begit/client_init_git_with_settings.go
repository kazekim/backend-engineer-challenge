package begit

//InitGitWithSettings init git object from settings
func (c *defaultClient) InitGitWithSettings(name, url string) Git {
	return &defaultGit{
		c:    c,
		name: name,
		url:  url,
	}
}
