package begit

type Git struct {
	c        Client
	name     string
	url      string
	path     *string
	fullPath *string
}
