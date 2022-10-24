package befiles

// SetMaxFileSize set max file size
func (c *defaultClient) SetMaxFileSize(size int64) {
	c.maxFileSize = size
}
