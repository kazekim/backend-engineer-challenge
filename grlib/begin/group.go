package begin

//Group create new RouterGroup from relativePath
func (c *defaultGin) Group(relativePath string) RouterGroup {
	g := c.g.Group(relativePath)
	return &defaultRouterGroup{
		g: *g,
	}
}
